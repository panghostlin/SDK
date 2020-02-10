################################################################################
## @Author:					Thomas Bouder <Tbouder>
## @Email:					Tbouder@protonmail.com
## @Date:					Sunday 05 January 2020 - 19:54:37
## @Filename:				makefile
##
## @Last modified by:		Tbouder
## @Last modified time:		Sunday 05 January 2020 - 19:55:03
################################################################################

all: build

build:
	@-echo "Generating Proto file"

	@-protoc --proto_path=./Protos --go_out=plugins=grpc,import_path=pictures:./Pictures Albums.proto
	@-protoc --proto_path=./Protos --go_out=plugins=grpc,import_path=pictures:./Pictures Pictures.proto
	@-protoc --proto_path=./Protos --go_out=plugins=grpc,import_path=keys:./Keys Keys.proto
	@-protoc --proto_path=./Protos --go_out=plugins=grpc,import_path=members:./Members Members.proto

clear:
	@-echo "Cleaning PB.GO"
	@-rm ./Keys/*.pb.go
	@-rm ./Pictures/*.pb.go
	@-rm ./Members/*.pb.go