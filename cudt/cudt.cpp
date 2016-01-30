
#include <cstdlib>
#include <cstring>
#include <netdb.h>
#include <signal.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <udt.h>
#include "cudt.h"


extern "C" {

  const int BACKLOG=1024;

  int startup(void) {
    return UDT::startup();
  }

  int cleanup() {
    return UDT::cleanup();
  }

  void set_error( const char* msg, struct udt_result* result ) {
    result->errorMsg = (char*)malloc(strlen(msg) + 1);
    strcpy(result->errorMsg, msg);
  }

  void udt_close( void* udtSocket) {
    UDT::close(*(UDTSOCKET*)udtSocket);
  }

  void udt_listen( const char* ipaddr, const char* port, struct udt_result** result ) {
    *result = (struct udt_result*)malloc(sizeof(udt_result));
    memset(*result, 0, sizeof(udt_result));

    UDTSOCKET sock;
    sock = UDT::socket(AF_INET, SOCK_STREAM, 0);

    sockaddr_in addr;
    addr.sin_family = AF_INET;
    addr.sin_port = htons(atoi(port));
    addr.sin_addr.s_addr = INADDR_ANY;
    memset(&(addr.sin_zero), '\0', 8);

    if( UDT::ERROR == UDT::bind(sock, (sockaddr*)&addr, sizeof(addr))) {
      set_error(UDT::getlasterror().getErrorMessage(), *result );
      return;
    }

    UDT::listen(sock, BACKLOG);

    (*result)->udtPointer = (void*)&sock;

    return;

  }

  void udt_accept(void* serv, struct udt_result** result ) {
    *result = (struct udt_result*)malloc(sizeof(udt_result));
    memset(*result, 0, sizeof(udt_result));
    
    int addrlen;
    sockaddr_in clientaddr;

    UDTSOCKET new_sock = UDT::accept(*(UDTSOCKET*)serv, (sockaddr*)&clientaddr, &addrlen);
    if(new_sock == UDT::INVALID_SOCK) {
      set_error(UDT::getlasterror().getErrorMessage(), *result);
      return;
    }

    const char* saddr = inet_ntoa(clientaddr.sin_addr);
    (*result)->addrString = (char*)malloc(strlen(saddr) + 1);
    strcpy((*result)->addrString, saddr);

    (*result)->udtPointer = (void*)&new_sock;
    return;

  }

}
