ws:
	mv generated ../generated && killall gopls && sleep 2 && mv ../generated ./generated

synth:
	cdktf synth

deploy:
	cdktf deploy