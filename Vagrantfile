# -*- mode: ruby -*-
# vi: set ft=ruby :

$ip = ENV['VAGRANT_IP'] || "172.22.22.22"

Vagrant.configure(2) do |config|
    arch = ""
    if ENV['VAGRANT_OS'] == "ubuntu" || !ENV['VAGRANT_OS']
        config.vm.box = "bento/ubuntu-14.04"
        arch = 'linux-amd64'
        home_dir = '/home'
    elsif ENV['VAGRANT_OS'] == "freebsd"
        config.vm.box = "freebsd/FreeBSD-11.0-CURRENT"
        arch = 'freebsd-amd64'
        home_dir = '/home'
    end
    
    config.vm.synced_folder ".", "/vagrant", disabled: true

    config.vm.network :private_network, ip: $ip

    config.vm.provision :shell, inline: "wget https://storage.googleapis.com/golang/go1.7.1.#{arch}.tar.gz"
    config.vm.provision :shell, inline: "tar xzvf ./go1.7.1.#{arch}.tar.gz"
    config.vm.provision :shell, inline: "mv ./go/ /usr/bin/go/"
    config.vm.provision :shell, inline: "echo 'export PATH=$PATH:/usr/bin/go/bin/' >> #{home_dir}/vagrant/.bashrc"
    config.vm.provision :shell, inline: "echo 'export GOROOT=/usr/bin/go/' >> #{home_dir}/vagrant/.bashrc"
    config.vm.provision :shell, inline: "echo 'export GOPATH=$GOPATH:#{home_dir}/vagrant/go/' >> #{home_dir}/vagrant/.bashrc"
    config.vm.provision :shell, inline: "mkdir -p #{home_dir}/vagrant/go/src/github.com/cpg1111/spawnd/ #{home_dir}/vagrant/go/bin/ #{home_dir}/vagrant/go/pkg/"

    config.vm.synced_folder "./", "#{home_dir}/vagrant/go/src/github.com/cpg1111/spawnd/", type: :nfs
end

