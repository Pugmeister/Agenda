const express = require('express');
const path = require('path');
const fs = require('fs');
const bodyParser = require('body-parser');
const multer  = require('multer');
const TelegramBot = require('node-telegram-bot-api');

const TELEGRAM_TOKEN = process.env.TELEGRAM_TOKEN;
if (!TELEGRAM_TOKEN) {
  console.error('Error: TELEGRAM_TOKEN is not set');
  process.exit(1);
}
const bot = new TelegramBot(TELEGRAM_TOKEN, { polling: false });

const app = express();
const PORT = process.env.PORT || 3000;

app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'templates'));
app.use(bodyParser.urlencoded({ extended: true }));
app.use('/static', express.static(path.join(__dirname, 'static')));

// Multer: сохраняем медиа в images/<project>/ под уникальным именем projectN.ext
const storage = multer.diskStorage({
  destination: (req, file, cb) => {
    const project = req.body.project.trim();
    const dir = path.join(__dirname, 'images', project);
    fs.mkdirSync(dir, { recursive: true });
    cb(null, dir);
  },
  filename: (req, file, cb) => {
    const project = req.body.project.trim();
    const ext = path.extname(file.originalname);
    const dir = path.join(__dirname, 'images', project);
    const existing = fs.readdirSync(dir)
      .map(f => {
        const m = f.match(new RegExp(`^${project}(\\d+)${ext}$`));
        return m ? parseInt(m[1], 10) : 0;
      });
    const next = existing.length ? Math.max(...existing) + 1 : 1;
    cb(null, `${project}${next}${ext}`);
  }
});
const upload = multer({ storage });

/** Главная — автоматическая серия постов */
app.get('/', (req, res) => {
  const imagesDir = path.join(__dirname, 'images');
  const projects = fs.existsSync(imagesDir)
    ? fs.readdirSync(imagesDir).filter(d => fs.statSync(path.join(imagesDir, d)).isDirectory())
    : [];
  res.render('index', {
    projects,
    postToValue: req.query.postTo || '',
    inPostValue: req.query.inPost || ''
  });
});

app.post('/send', async (req, res) => {
  // Сохраняем ссылки
  if (req.body.postTo) {
    fs.writeFileSync(path.join(__dirname, 'links/postTo/inp.html'), req.body.postTo.trim() + '\n', 'utf-8');
  }
  if (req.body.inPost) {
    fs.writeFileSync(path.join(__dirname, 'links/inPost/inp.html'), req.body.inPost.trim() + '\n', 'utf-8');
  }

  const project   = req.body.project.trim();
  const countScen = parseInt(req.body.count, 10);
  const chatId    = parseInt(req.body.chat, 10);
  const imgtype   = req.body.imgtype;

  const postToLinks = (req.body.postTo || '')
    .replace(/\r/g, '').split('\n').map(l => l.trim()).filter(l => l);
  const inPostLinks = (req.body.inPost || '')
    .replace(/\r/g, '').split('\n').map(l => l.trim()).filter(l => l);

  try {
    await sendPosts(project, countScen, chatId, imgtype, postToLinks, inPostLinks);
    const q = new URLSearchParams({ postTo: req.body.postTo||'', inPost: req.body.inPost||'' }).toString();
    res.redirect('/?' + q);
  } catch (err) {
    console.error(err);
    res.status(500).send(`Ошибка при отправке: ${err.message}`);
  }
});

/** Страница сохранения в папки — без Telegram */
app.get('/create', (req, res) => {
  res.render('create');
});

app.post('/create', upload.single('media'), (req, res) => {
  try {
    const project = req.body.project.trim();
    const caption = req.body.caption || '';

    // Сохраняем сценарий как scenarios/<project><N>.html
    const scenDir = path.join(__dirname, 'scenarios');
    fs.mkdirSync(scenDir, { recursive: true });
    const ext = '.html';
    const existing = fs.readdirSync(scenDir)
      .map(f => {
        const m = f.match(new RegExp(`^${project}(\\d+)${ext}$`));
        return m ? parseInt(m[1], 10) : 0;
      });
    const next = existing.length ? Math.max(...existing) + 1 : 1;
    const scenPath = path.join(scenDir, `${project}${next}${ext}`);
    const htmlContent = `<div class="post">\n${caption}\n</div>`;
    fs.writeFileSync(scenPath, htmlContent, 'utf-8');

    res.send(`
      Сохранено:<br/>
      • Медиа: <strong>${req.file.filename}</strong><br/>
      • Сценарий: <strong>${project}${next}${ext}</strong><br/>
      <p><a href="/create">Сохранить ещё</a> | <a href="/">К авто-серии</a></p>
    `);
  } catch (err) {
    console.error(err);
    res.status(500).send('Ошибка при сохранении: ' + err.message);
  }
});

/** Отправка серии постов в Telegram */
async function sendPosts(project, countScen, chatId, imgtype, postToLinks, inPostLinks) {
  const num = Math.min(postToLinks.length, inPostLinks.length);
  for (let k = 0; k < num; k++) {
    const scenIdx = (k % countScen) + 1;
    let scenPath = path.join(__dirname, 'scenarios', `${project}${scenIdx}.html`);
    if (!fs.existsSync(scenPath)) {
      // fallback на первый сценарий, если специфический отсутствует
      scenPath = path.join(__dirname, 'scenarios', `${project}1.html`);
    }
    // Удаляем div-теги перед отправкой
    let html = fs.readFileSync(scenPath, 'utf-8')
      .replace(/<div[^>]*>/g, '')
      .replace(/<\/div>/g, '')
      .trim();

    // Отправка заголовка
    await bot.sendMessage(chatId, `👇 пост для ${postToLinks[k]}`);

    // Отправка медиа
    const media = path.join(__dirname, 'images', project, `${project}${scenIdx}.${imgtype}`);
    if (fs.existsSync(media)) {
      if (imgtype === 'mp4') {
        await bot.sendVideo(chatId, media, { caption: html, parse_mode: 'HTML' });
      } else {
        await bot.sendPhoto(chatId, media, { caption: html, parse_mode: 'HTML' });
      }
    }
  }
}

app.listen(PORT, () => console.log(`Server running on http://localhost:${PORT}`));