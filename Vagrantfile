# -*- mode: ruby -*-
# vi: set ft=ruby :

$ip = ENV['VAGRANT_IP'] || "172.22.22.22"

Vagrant.configure(2) do |config|
    arch = ""
    if ENV['VAGRANT_OS'] == "ubuntu" || !ENV['VAGRANT_OS']
        config.vm.box = "bento/ubuntu-14.04"
        arch = 'linux-amd64'
        home_dir = '/home'
        pkg_mgr = 'apt-get'
    elsif ENV['VAGRANT_OS'] == "freebsd"
        config.vm.box = "bento/freebsd-10.2"
        arch = 'freebsd-amd64'
        home_dir = '/home'
        pkg_mgr = 'pkg'
    end
    
    config.vm.synced_folder ".", "/vagrant", disabled: true

    config.vm.network :private_network, ip: $ip

    config.vm.provision :shell, inline: "#{pkg_mgr} update && #{pkg_magr} install -y curl"
    config.vm.provision :shell, inline: "curl -o ./go1.7.1.#{arch}.tar.gz https://storage.googleapis.com/golang/go1.7.1.#{arch}.tar.gz"
    config.vm.provision :shell, inline: "tar xzvf ./go1.7.1.#{arch}.tar.gz"
    config.vm.provision :shell, inline: "mv ./go/ /usr/local/go/"
    config.vm.provision :shell, inline: "echo 'export PATH=$PATH:/usr/local/go/bin/' >> #{home_dir}/vagrant/.bashrc"
    config.vm.provision :shell, inline: "echo 'export GOROOT=/usr/local/go/' >> #{home_dir}/vagrant/.bashrc"
    config.vm.provision :shell, inline: "echo 'export GOPATH=#{home_dir}/vagrant/go/' >> #{home_dir}/vagrant/.bashrc"
    config.vm.provision :shell, inline: "echo 'export PATH=$PATH:$GOPATH/bin/' >> #{home_dir}/vagrant/.bashrc"
    config.vm.provision :shell, inline: "mkdir -p #{home_dir}/vagrant/go/src/github.com/cpg1111/spawnd/ #{home_dir}/vagrant/go/bin/ #{home_dir}/vagrant/go/pkg/"
    config.vm.provision :shell, inline: "chown vagrant:vagrant #{home_dir}/vagrant/go/bin/ #{home_dir}/vagrant/go/pkg/ #{home_dir}/vagrant/go/src/"

    config.vm.synced_folder "./", "#{home_dir}/vagrant/go/src/github.com/cpg1111/spawnd/", type: :nfs
end

