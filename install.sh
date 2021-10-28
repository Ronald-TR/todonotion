curl -L -o ${HOME}/todonotion-linux-386.tar.gz https://github.com/Ronald-TR/todonotion/releases/download/v0.1.0/todonotion-linux-386.tar.gz
tar -xzvf ~/todonotion-linux-386.tar.gz -C ${HOME}
chmod +x ~/todonotion
sudo ln -s ~/todonotion /usr/local/bin
