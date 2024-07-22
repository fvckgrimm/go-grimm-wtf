// Store command history
let commandHistory = [];
let historyIndex = -1;

// List of available commands
const availableCommands = [
  'help', 'clear', 'frens', 'contact', 'git', "games", 'bookmarks', 'blog', 'changetheme', 'theme', 'search', 'about'
];

const availableThemes = [
  'mocha', 'frappe', 'tokyo-night', 'dracula', 'gruvbox-dark'
];

function refreshTime() {
  const timeDisplay = document.getElementById("time");
  const dateString = new Date().toLocaleString();
  const formattedString = dateString.replace(", ", " - ");
  timeDisplay.textContent = formattedString;
}
setInterval(refreshTime, 1000);

function changeTheme(theme) {
  if (availableThemes.includes(theme)) {
    document.body.classList.remove(...availableThemes);
    document.body.classList.add(theme);

    // Ensure the font-mono class is always present
    if (!document.body.classList.contains('font-mono')) {
      document.body.classList.add('font-mono');
    }

    localStorage.setItem('theme', theme);

    //document.body.className = `font-mono ${theme}`;
    return `Theme changed to ${theme}`;
  } else {
    return `Invalid theme option. Available options: [${availableThemes.join(' | ')}]`;
  }
}

function loadTheme() {
  const savedTheme = localStorage.getItem('theme');
  if (savedTheme && availableThemes.includes(savedTheme)) {
    document.body.classList.add(savedTheme);
  } else {
    document.body.classList.add('mocha'); // Default theme
  }
}

function search(provider, query) {
  const searchUrls = {
    brave: 'https://search.brave.com/search?q=',
    google: 'https://www.google.com/search?q=',
    youtube: 'https://www.youtube.com/results?search_query=',
    ddg: 'https://duckduckgo.com/?q='
  };

  if (searchUrls[provider]) {
    window.open(searchUrls[provider] + encodeURIComponent(query), '_blank');
    return `Searching ${provider} for "${query}"`;
  } else {
    return `Invalid search provider. Available options: [${Object.keys(searchUrls).join(' | ')}]`;
  }
}

function runCommand(event) {
  event.preventDefault();
  const input = document.getElementById("command-input").value.trim();
  const output = document.querySelector(".terminal-output");
  const p = document.createElement("p");
  const [command, ...args] = input.split(' ');

  // Add command to history
  commandHistory.unshift(input);
  historyIndex = -1;

  p.textContent = `user@grimm.wtf$ ${input}`;
  output.appendChild(p);
  const response = document.createElement("p");


  if (args[0] === '-h'){
      response.textContent = getHelpText(command);
  } else {
    switch (command) {
      case "help":
        response.textContent = `Available commands: ${availableCommands.join(', ')}`;
        break;
      case "changetheme":
      case "theme":
        if (args.length === 0) {
            response.textContent = `Please provide a theme option. Available options: [${availableThemes.join(' | ')}]`;
        } else {
            response.textContent = changeTheme(args[0]);
        }
        break;
      case "frens":
        window.open('https://degen.wtf')
        response.textContent = "user@grimm.wtf$ frens: opening degen.wtf...";
        break;
      case "contact":
        response.innerHTML = `
          email: <a href="mailto:its@grimm.wtf" title="its@grimm.wtf">its@grimm.wtf</a><br>
          github: <a href="https://github.com/fvckgrimm" title="https://github.com/fvckgrimm">https://github.com/fvckgrimm</a><br>
          discord: @grimm.wtf<br>
          x.com: <a href="https://x.com/grimmdusk" title="https://x.com/grimmdusk">@grimmdusk</a><br>
          bluesky: <a href="https://bsky.app/profile/grimm.wtf" title="https://bsky.app/profile/grimm.wtf">@grimm.wtf</a>`;
        break;
      case "about":
            response.innerHTML = `
            <p>I'm Grimm, a fake schizo privacy enthusiast, ai script maker, and wannabe hackermanz. You have no reason to be here, please leave ᗜˬᗜ </p>
            `
            break;
      case "git":
        window.open('https://git.grimm.wtf')
        response.textContent = "user@grimm.wtf$ git: opening git instance...";
        break;
      case "bookmarks":
        response.innerHTML = `
          /g/sec/urity: <a href="https://pastebin.com/UY7RxEqp" title="/g/sec/"> https://pastebin.com/UY7RxEqp </a><br>
          Crow's Nest: <a href="https://www.crow.rip/crows-nest" title="Crows Nest"> https://www.crow.rip/crows-nest </a><br>
          revshells: <a href="https://www.revshells.com/" title="revshells"> https://www.revshells.com/ </a><br>
          TryHackMe: <a href="https://tryhackme.com/" title="tryhackme"> https://tryhackme.com/ </a><br>
          HackTheBox: <a href="https://app.hackthebox.com/" title="HackTheBox"> https://app.hackthebox.com/ </a>`;
        break;
      case "games":
        response.innerHTML = `
          MonkeyType: <a href="https://monkeytype.com" title="MonkeyType"> https://monkeytype.com </a><br>
          Chess: <a href="https://chess.com" title="Chess"> https://chess.com </a><br>
          Tetris: <a href="https://tetr.io" title="tetris"> https://tetr.io </a>
            `;
        break;
      case "clear":
        clearTerminal();
        return;
      case "blog":
        p.textContent = "user@grimm.wtf$ Recent blog posts:";
        output.appendChild(p);
        fetch('/api/recent-posts')
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => {
                if (data.success && data.response && data.response.length > 0) {
                    data.response.forEach(post => {
                        const postElement = document.createElement("p");
                        postElement.innerHTML = `- <a href="/blog/${post.slug}">${post.title}</a>`;
                        output.appendChild(postElement);
                    });
                } else {
                    const noPostsElement = document.createElement("p");
                    noPostsElement.textContent = "No recent posts found.";
                    output.appendChild(noPostsElement);
                }
                const allPostsLink = document.createElement("p");
                allPostsLink.innerHTML = '<a href="/blog">View all posts</a>';
                output.appendChild(allPostsLink);
            })
            .catch(error => {
                console.error("Error fetching blog posts:", error);
                const errorElement = document.createElement("p");
                errorElement.textContent = `Failed to load recent blog posts: ${error.message}`;
                output.appendChild(errorElement);
            });
        break;
      case "search":
        if (args.length < 2) {
          response.textContent = "Usage: search [brave|google|youtube|ddg] <search query>";
        } else {
          const [provider, ...queryParts] = args;
          response.textContent = search(provider, queryParts.join(' '));
        }
        break;
      default:
        response.textContent = "Command not recognized. Type 'help' for a list of available commands.";
        break;
    }
  }

  output.appendChild(response);
  document.getElementById("command-input").value = "";
  output.scrollTop = output.scrollHeight;
}

function getHelpText(command) {
  switch (command) {
    case "help":
      return "Displays a list of available commands.";
    case "clear":
      return "Clears the terminal output.";
    case "frens":
      return "Opens the degen.wtf website.";
    case "contact":
      return "Displays contact information.";
    case "git":
      return "Opens the git instance.";
    case "bookmarks":
      return "Lists of resources.";
    case "blog":
      return "Displays recent blog posts.";
    case "changetheme":
    case "theme":
      return `Changes the terminal theme. Usage: changetheme [${availableThemes.join(' | ')}]`;
    case "search":
      return "Searches using specified provider. Usage: search [brave|google|youtube|ddg] <search query>";
    default:
      return "No help available for this command.";
  }
}

function clearTerminal() {
  const output = document.querySelector(".terminal-output");
  while (output.children.length > 2) {
    output.removeChild(output.lastChild);
  }
  document.getElementById("command-input").value = "";
}


// Add tab completion function
function tabComplete(input) {
  const [partialCommand, ...args] = input.split(' ');
  
  if (args.length === 0) {
    // Complete command
    const matches = availableCommands.filter(cmd => cmd.startsWith(partialCommand));
    if (matches.length === 1) {
      return matches[0];
    } else if (matches.length > 1) {
      console.log(matches.join('  '));
      return partialCommand;
    }
  } else if ((partialCommand === 'changetheme' || partialCommand === 'theme') && args.length === 1) {
    // Complete theme options
    const partialTheme = args[0];
    const matches = availableThemes.filter(theme => theme.startsWith(partialTheme));
    if (matches.length === 1) {
      return `${partialCommand} ${matches[0]}`;
    } else if (matches.length > 1) {
      console.log(matches.join('  '));
      return input;
    }
  }
  
  return input;
}


document.addEventListener("keydown", function(event) {
  if (event.ctrlKey && event.key === 'l') {
    event.preventDefault();
    clearTerminal();
  } else if (event.key === "ArrowUp") {
    event.preventDefault();
    if (historyIndex < commandHistory.length - 1) {
      historyIndex++;
      document.getElementById("command-input").value = commandHistory[historyIndex];
    }
  } else if (event.key === "ArrowDown") {
    event.preventDefault();
    if (historyIndex > 0) {
      historyIndex--;
      document.getElementById("command-input").value = commandHistory[historyIndex];
    } else if (historyIndex === 0) {
      historyIndex = -1;
      document.getElementById("command-input").value = "";
    }
  }
});

document.getElementById("command-input").addEventListener("keydown", function(event) {
  if (event.key === "Enter") {
    runCommand(event);
  } else if (event.key === "Tab") {
    event.preventDefault();
    this.value = tabComplete(this.value);
  }
});

document.addEventListener("DOMContentLoaded", loadTheme);
