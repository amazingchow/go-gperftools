typedef void* HeapLeakCheckerCGO;

#ifdef __cplusplus
extern "C" {
#endif

HeapLeakCheckerCGO start(char* name);
void stop(HeapLeakCheckerCGO handler);

#ifdef __cplusplus
}
#endif
