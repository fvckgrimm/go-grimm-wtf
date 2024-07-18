const availableThemes = [
            'mocha', 'frappe', 'tokyo-night', 'dracula', 'gruvbox-dark'
        ];

function loadTheme() {
  const savedTheme = localStorage.getItem('theme');
  if (savedTheme && availableThemes.includes(savedTheme)) {
    document.body.classList.add(savedTheme);
  } else {
    document.body.classList.add('mocha'); // Default theme
  }
}

document.addEventListener("DOMContentLoaded", loadTheme);

