// Created by cgo -cdefs - DO NOT EDIT
// cgo -cdefs defs_solaris.go defs_solaris_amd64.go

typedef struct SemT SemT;
typedef struct Fpregset Fpregset;
typedef struct PortEvent PortEvent;
typedef struct PthreadAttr PthreadAttr;
typedef struct Stat Stat;

#pragma pack on

struct SemT {
	uint32	sem_count;
	uint16	sem_type;
	uint16	sem_magic;
	uint64	sem_pad1[3];
	uint64	sem_pad2[2];
};

struct Fpregset {
	byte	fp_reg_set[528];
};

struct PortEvent {
	int32	portev_events;
	uint16	portev_source;
	uint16	portev_pad;
	uint64	portev_object;
	byte	*portev_user;
};
typedef	uint32	Pthread;
struct PthreadAttr {
	byte	*__pthread_attrp;
};

struct Stat {
	uint64	st_dev;
	uint64	st_ino;
	uint32	st_mode;
	uint32	st_nlink;
	uint32	st_uid;
	uint32	st_gid;
	uint64	st_rdev;
	int64	st_size;
	Timespec	st_atim;
	Timespec	st_mtim;
	Timespec	st_ctim;
	int32	st_blksize;
	byte	Pad_cgo_0[4];
	int64	st_blocks;
	int8	st_fstype[16];
};

#pragma pack off
