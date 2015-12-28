#!/bin/sh

echo 'Preparing to build UDT C library.'

if [ -z ${GOPATH} ]; then
  echo 'You must define your GOPATH environment variable prior to running this script.'
  exit 1
else
  echo "GOPATH present at ${GOPATH}"
  echo "Checking architecture."
fi

os=''
arch=''

case "$(uname -s)" in
  'Darwin')
  os='OSX'
  ;;
  'Linux')
  os='LINUX'
  ;;
  *)
  echo 'Unknown OS.'
  exit 1
  ;;
esac

case "$(uname -m)" in
  'x86_64')
  arch='AMD64'
  ;;
  *)
  echo 'Unknown architecture.'
  exit 1
  ;;
esac



echo "Building for ${os} ${arch}"

script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
udt4_project_dir="${script_dir}/vendor/git.code/sf.net/udt/udt4"

if [ ! -d "${udt4_project_dir}" ]; then
  echo "Missing source code for UDT4 C library. Try git submodule update --recursive"
  exit 1
else
  echo "UDT4 C source code is present."
fi

current=$(pwd)

cd "${udt4_project_dir}"

make clean
make -e os=$os arch=$arch


if [ $? -ne 0 ]; then
  echo "Build failed with error ${?}"
  exit 1
else
  echo "UDT4 build succeeded."
fi

app/test

if [ $? -ne 0 ]; then
  echo "UDT4 test failed."
  exit 1
else
  echo "UDT4 tests succeeded."
fi


ln -s ${udt4_project_dir} "${GOPATH}/."

libpathname='LD_LIBRARY_PATH'

if [ "$os" = 'OSX' ]; then
  libpathname='DYLD_LIBRARY_PATH'
fi


echo "UDT4 build was successful."
echo "Set ${libpathname}=${GOPATH}/udt4/src"
echo "UDT4 header files are in ${GOPATH}/udt4/src"




#cd vendor/git.code/sf.net/udt/udt4
