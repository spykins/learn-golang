# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
$script = <<-SCRIPT
    sudo apt -y update
    sudo apt -y install build-essential
    sudo apt-get -y install manpages-dev
    # sudo apt-get -y install clang-10
    sudo apt-get -y install wget
# installing python7
    sudo add-apt-repository ppa:deadsnakes/ppa   
    sudo apt-get update   
# installing golang
    wget -O /tmp/go1.17.linux-amd64.tar.gz https://golang.org/dl/go1.17.linux-amd64.tar.gz /tmp
    rm -rf /usr/local/go && tar -C /usr/local -xzf /tmp/go1.17.linux-amd64.tar.gz
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /home/ubuntu/.bashrc
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /home/vagrant/.bashrc

SCRIPT
Vagrant.configure("2") do |config|

    config.vm.box = "ubuntu/trusty64"
    config.vm.network "public_network", bridge: "en3: Wi-Fi (AirPort)"
    config.vm.synced_folder "golang_code/", "/srv/golang_code" 
    config.vm.provision "shell", inline: $script
end

