# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "ubuntu/xenial64"
  config.vm.hostname = "testland"
  config.vm.synced_folder ENV['GOPATH'], "/go"

  # Golang
  config.vm.provision "shell", inline: "sudo wget -q -O /tmp/go.tar.gz https://dl.google.com/go/go1.9.3.linux-amd64.tar.gz"
  config.vm.provision "shell", inline: "sudo tar -C /usr/local -xzf /tmp/go.tar.gz"
  config.vm.provision "shell", inline: "echo 'export PATH=$PATH:/usr/local/go/bin' | sudo tee /etc/profile.d/go.sh"

  # Docker
  config.vm.provision "shell", inline: "curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -"
  config.vm.provision "shell", inline: "sudo add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\""
  config.vm.provision "shell", inline: "sudo apt-get update"
  config.vm.provision "shell", inline: "apt-cache policy docker-ce"
  config.vm.provision "shell", inline: "sudo apt-get install -y docker-ce"
  config.vm.provision "shell", inline: "sudo usermod -aG docker ubuntu"

end
