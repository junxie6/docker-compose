# .bash_profile

# Get the aliases and functions
if [ -f ~/.bashrc ]; then
        . ~/.bashrc
fi

# User specific environment and startup programs

export PATH="$PATH:/usr/local/bin:/usr/local/go/bin:$HOME/.local/bin:$HOME/bin:$HOME/go/bin"