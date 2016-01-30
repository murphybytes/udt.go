#ifndef __CUDT_H__
#define __CUDT_H__


#ifdef __cplusplus
extern "C" {
#endif


struct udt_result {
  void* udtPointer;
  char* errorMsg;
  char* addrString;
} ;

int startup(void);
int cleanup(void);
void udt_listen( const char*, const char*, struct udt_result** );
void udt_close(void*);
void udt_accept(void* , struct udt_result**);


#ifdef __cplusplus
}
#endif

#endif
