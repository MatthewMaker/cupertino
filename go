gem build cupertino.gemspec
gem install cupertino-0.8.1.gem
#ln -s ./bin/ios .
say "done."
export SSL_CERT_FILE=charles-proxy-ssl-proxying-certificate.crt
