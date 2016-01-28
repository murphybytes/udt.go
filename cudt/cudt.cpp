
#include <cstdlib>
#include <cstring>
#include <netdb.h>
#include <signal.h>
#include <unistd.h>
#include <udt.h>
#include "cudt.h"


extern "C" {



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

    addrinfo hints;
    addrinfo* res;

    memset(&hints, 0, sizeof(struct addrinfo));
    hints.ai_flags = AI_PASSIVE;
    hints.ai_family = AF_INET;
    hints.ai_socktype = SOCK_STREAM;

    if( 0 != getaddrinfo(strlen(ipaddr) > 0 ? ipaddr : NULL, port, &hints, &res )) {
      set_error("Illegal port number or port is busy", *result);
      return;
    }

    UDTSOCKET sock;
    sock = UDT::socket(res->ai_family, res->ai_socktype, res->ai_protocol);

    if( UDT::ERROR == UDT::bind(sock, res->ai_addr, res->ai_addrlen)) {
      set_error(UDT::getlasterror().getErrorMessage(), *result );
      return;
    }



    (*result)->udtPointer = (void*)&sock;
    freeaddrinfo(res);
    return;

    // const char* error = "Something broken";
    // result->errorMsg = (char*)malloc(strlen(error) + 1);
    // strcpy(result->errorMsg, error );

  }

}
