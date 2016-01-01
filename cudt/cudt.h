#ifndef __CUDT_H__
#define __CUDT_H__


#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
  void* pUDTSOCKET;
} udt_connection;

int startup(void);
int cleanup(void);

#ifdef __cplusplus
}
#endif

#endif
