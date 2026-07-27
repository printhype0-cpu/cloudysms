const fs = require('fs');
const path = require('path');

const directoryPath = path.join(__dirname, 'frontend/src');

const routesToUpdate = ["chat", "chatbot", "analytics", "templates", "flows", "campaigns", "calling", "settings", "profile", "dashboard"];

function updateFilesInDirectory(dir) {
  const files = fs.readdirSync(dir);

  for (const file of files) {
    const fullPath = path.join(dir, file);
    const stat = fs.statSync(fullPath);

    if (stat.isDirectory()) {
      updateFilesInDirectory(fullPath);
    } else if (file.endsWith('.vue') || file.endsWith('.ts')) {
      if (fullPath.includes('router\\index.ts') || fullPath.includes('router/index.ts')) {
        continue;
      }
      
      let content = fs.readFileSync(fullPath, 'utf8');
      let modified = false;

      for (const route of routesToUpdate) {
        // to="/route"
        const regex1 = new RegExp(`to="/${route}`, 'g');
        if (regex1.test(content)) {
          content = content.replace(regex1, `to="/app/${route}`);
          modified = true;
        }

        // push('/route'
        const regex2 = new RegExp(`push\\('/${route}`, 'g');
        if (regex2.test(content)) {
          content = content.replace(regex2, `push('/app/${route}`);
          modified = true;
        }

        // replace('/route'
        const regex3 = new RegExp(`replace\\('/${route}`, 'g');
        if (regex3.test(content)) {
          content = content.replace(regex3, `replace('/app/${route}`);
          modified = true;
        }

        // redirect: '/route'
        const regex4 = new RegExp(`redirect:\\s*['"]/${route}`, 'g');
        if (regex4.test(content)) {
          content = content.replace(regex4, `redirect: '/app/${route}`);
          modified = true;
        }
      }

      // Special case for AppLayout brand link
      if (fullPath.includes('AppLayout.vue')) {
        if (content.includes('to="/"')) {
           content = content.replace(/to="\/"/g, 'to="/app/dashboard"');
           modified = true;
        }
      }

      if (modified) {
        fs.writeFileSync(fullPath, content, 'utf8');
        console.log(`Updated: ${fullPath}`);
      }
    }
  }
}

updateFilesInDirectory(directoryPath);
console.log("Link update complete!");
