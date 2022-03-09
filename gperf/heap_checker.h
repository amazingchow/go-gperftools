typedef void* HeapLeakCheckerCGO;

#ifdef __cplusplus
extern "C" {
#endif

HeapLeakCheckerCGO start(char* name);
void stop(HeapLeakCheckerCGO i);

#ifdef __cplusplus
}
#endif
