
#include <cstdlib>
#include <cstring>
#include <netdb.h>
#include <signal.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <udt.h>
#include "cudt.h"
#include <iostream>

using namespace std;


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

  void udt_close( int sock) {
    UDT::close(sock);
  }

  void udt_connect( const char* ipaddr, const char* port, struct udt_result** result ) {
    cout << "called connect with " << ipaddr << " " << port << endl;
    *result = (struct udt_result*)malloc(sizeof(udt_result));
    memset(*result, 0, sizeof(udt_result));

    UDTSOCKET sock;
    sock = UDT::socket(AF_INET, SOCK_STREAM, 0);

    sockaddr_in serv_addr;
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(atoi(port));
    inet_pton(AF_INET, ipaddr, &serv_addr.sin_addr);

    memset(&(serv_addr.sin_zero), '\0', 8);

    // connect to the server, implict bind
    if (UDT::ERROR == UDT::connect(sock, (sockaddr*)&serv_addr, sizeof(serv_addr))) {
      set_error( UDT::getlasterror().getErrorMessage(), *result );
      return;
    }

    (*result)->udtSocket = sock;

    cout << "connect returns" << endl;

    return;

  }

  void udt_listen( const char* ipaddr, const char* port, struct udt_result** result ) {
    cout << "called listen with " << ipaddr << " " << port << endl;

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
    cout << "post listen" << endl;
    (*result)->udtSocket = sock;
    cout << "listen returned" << endl;
    return;

  }

  void udt_accept(int serv, struct udt_result** result ) {
    cout << "Called accept" << endl;
    *result = (struct udt_result*)malloc(sizeof(udt_result));
    memset(*result, 0, sizeof(udt_result));

    int addrlen;
    sockaddr_in clientaddr;

    UDTSOCKET new_sock = UDT::accept(serv, (sockaddr*)&clientaddr, &addrlen);
    if(new_sock == UDT::INVALID_SOCK) {
      set_error(UDT::getlasterror().getErrorMessage(), *result);
      return;
    }

    const char* saddr = inet_ntoa(clientaddr.sin_addr);
    (*result)->addrString = (char*)malloc(strlen(saddr) + 1);
    strcpy((*result)->addrString, saddr);

    (*result)->udtSocket = new_sock;
    cout << "Accept returns" << endl;
    return;

  }

}
