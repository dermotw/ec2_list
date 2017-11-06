world:
	go build -ldflags "-s -w" 
	upx ec2_list

install:
	cp ec2_list ${HOME}/.config/argos/ec2-go.10s.cmd

uninstall:
	rm ${HOME}/.config/argos/ec2-go.10s.cmd

clean:
	rm ec2_list
