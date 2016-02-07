
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
  const int RCVBUFFERSIZE=100000;

  int startup(void) {
    return UDT::startup();
  }

  int cleanup() {
    return UDT::cleanup();
  }

  void set_error( struct udt_result* result ) {
    const char* msg = UDT::getlasterror().getErrorMessage();
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
      set_error( *result );
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
      set_error( *result );
      return;
    }

    UDT::listen(sock, BACKLOG);
    (*result)->udtSocket = sock;
    return;

  }

  void udt_accept(int serv, struct udt_result** result ) {
    cout << "Called accept server socket -> " << serv << endl;
    *result = (struct udt_result*)malloc(sizeof(udt_result));
    memset(*result, 0, sizeof(udt_result));

    int addrlen;
    sockaddr_in clientaddr;

    UDTSOCKET new_sock = UDT::accept(serv, (sockaddr*)&clientaddr, &addrlen);
    if(new_sock == UDT::INVALID_SOCK) {
      set_error(*result);
      return;
    }

    const char* saddr = inet_ntoa(clientaddr.sin_addr);
    (*result)->addrString = (char*)malloc(strlen(saddr) + 1);
    strcpy((*result)->addrString, saddr);

    (*result)->udtSocket = new_sock;
    return;

  }

  void udt_send( int sock, const char* buffer, int buff_cap, int* bytes_sent, struct udt_result** result ) {

    *result = (struct udt_result*)malloc(sizeof(udt_result));
    memset(*result, 0, sizeof(udt_result));
    *bytes_sent = 0;

    int this_send = 0;
    int sent = 0;


    while( sent < buff_cap ) {
      this_send = UDT::send(sock, buffer + sent, buff_cap - sent, 0 );

      if(UDT::ERROR == this_send ) {
        set_error(*result);
        return;
      }

      sent += this_send;
    }

    *bytes_sent = sent;

  }

  void udt_recv( int sock, char* buffer, int buff_cap, int* bytes_read, struct
    udt_result** result ) { *result = (struct udt_result*)malloc(sizeof(udt_result));

    memset(*result, 0, sizeof(udt_result));

    *bytes_read = 0;

    if( UDT::ERROR == (*bytes_read = UDT::recv(sock, buffer, buff_cap, 0 ))){
      set_error(*result);
      return;
    }


  }
}
