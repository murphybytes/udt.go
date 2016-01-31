#ifndef __CUDT_H__
#define __CUDT_H__


#ifdef __cplusplus
extern "C" {
#endif


struct udt_result {
  int udtSocket;
  char* errorMsg;
  char* addrString;
} ;

int startup(void);
int cleanup(void);
void udt_listen( const char*, const char*, struct udt_result** );
void udt_close(int);
void udt_accept(int , struct udt_result**);
void udt_connect( const char*, const char*, struct udt_result** );



#ifdef __cplusplus
}
#endif

#endif
