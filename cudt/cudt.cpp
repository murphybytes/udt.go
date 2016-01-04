#include <udt.h>
#include "cudt.h"

extern "C" {



  int startup(void) {
    return UDT::startup();
  }

  int cleanup() {
    return UDT::cleanup();
  }

  void udt_listen( const char* ipaddr, const char* port, struct udt_result* result ) {
    const char* error = "Something broken";
    result->errorMsg = (char*)malloc(strlen(error) + 1);
    strcpy(result->errorMsg, error );

  }

}
