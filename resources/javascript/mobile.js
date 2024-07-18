// Check if there is a saved theme in local storage
var savedTheme = localStorage.getItem('theme');
if (savedTheme) {
    setTheme(savedTheme); // Apply saved theme
}

document.getElementById('asciicat').addEventListener('pointerdown', function(event) {
    var dropdown = document.getElementById('theme-dropdown');
    dropdown.style.display = dropdown.style.display === 'none' ? 'block' : 'none';
    event.target.setPointerCapture(event.pointerId);
});

document.getElementById('asciicat').addEventListener('pointerup', function(event) {
    event.target.releasePointerCapture(event.pointerId);
});

// Handle theme selection
var themeLinks = document.querySelectorAll('#theme-dropdown a');
themeLinks.forEach(function(link) {
    link.addEventListener('click', function(event) {
        event.preventDefault();
        var themeName = this.getAttribute('data-theme');
        setTheme(themeName);
        localStorage.setItem('theme', themeName); // Save selected theme to local storage
        closeDropdown();
    });
});

function setTheme(themeName) {
    var body = document.body;
    body.className = ''; // Clear existing classes
    body.classList.add('bg-custom', 'text-custom', themeName);
}

function closeDropdown() {
    document.getElementById('theme-dropdown').style.display = 'none';
}

// Event listener to close the dropdown when clicking outside of it
document.addEventListener('click', function(event) {
    var dropdown = document.getElementById('theme-dropdown');
    var asciicat = document.getElementById('asciicat');

    if (dropdown.style.display === 'block' && !dropdown.contains(event.target) && event.target !== asciicat) {
        closeDropdown();
    }
});

document.addEventListener('touchstart', function(event) {
    var dropdown = document.getElementById('theme-dropdown');
    var asciicat = document.getElementById('asciicat');

    if (dropdown.style.display === 'block' && !dropdown.contains(event.target) && event.target !== asciicat) {
        closeDropdown();
    }
});
