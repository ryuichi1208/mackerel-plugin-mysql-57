package mpmysql

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphDefinition_DisableInnoDB(t *testing.T) {
	var mysql MySQLPlugin

	mysql.DisableInnoDB = true
	graphdef := mysql.GraphDefinition()
	if len(graphdef) != 8 {
		t.Errorf("GetTempfilename: %d should be 7", len(graphdef))
	}
}

func TestGraphDefinition(t *testing.T) {
	var mysql MySQLPlugin

	graphdef := mysql.GraphDefinition()
	if len(graphdef) != 29 {
		t.Errorf("GetTempfilename: %d should be 28", len(graphdef))
	}
}

func TestGraphDefinition_DisableInnoDB_EnableExtended(t *testing.T) {
	var mysql MySQLPlugin

	mysql.DisableInnoDB = true
	mysql.EnableExtended = true
	graphdef := mysql.GraphDefinition()
	if len(graphdef) != 18 {
		t.Errorf("GetTempfilename: %d should be 18", len(graphdef))
	}
}

func TestGraphDefinition_EnableExtended(t *testing.T) {
	var mysql MySQLPlugin

	mysql.EnableExtended = true
	graphdef := mysql.GraphDefinition()
	if len(graphdef) != 39 {
		t.Errorf("GetTempfilename: %d should be 39", len(graphdef))
	}
}

func TestParseProcStat56(t *testing.T) {
	stub := `=====================================
2015-03-09 20:11:22 7f6c0c845700 INNODB MONITOR OUTPUT
=====================================
Per second averages calculated from the last 6 seconds
-----------------
BACKGROUND THREAD
-----------------
srv_master_thread loops: 178 srv_active, 0 srv_shutdown, 1244368 srv_idle
srv_master_thread log flush and writes: 1244546
----------
SEMAPHORES
----------
OS WAIT ARRAY INFO: reservation count 227
OS WAIT ARRAY INFO: signal count 220
Mutex spin waits 923, rounds 9442, OS waits 193
RW-shared spins 19, rounds 538, OS waits 16
RW-excl spins 5, rounds 476, OS waits 13
Spin rounds per wait: 10.23 mutex, 28.32 RW-shared, 95.20 RW-excl
------------
TRANSACTIONS
------------
Trx id counter 1093821584
Purge done for trx's n:o < 1093815563 undo n:o < 0 state: running but idle
History list length 649
LIST OF TRANSACTIONS FOR EACH SESSION:
---TRANSACTION 0, not started
MySQL thread id 27954, OS thread handle 0x7f6c0c845700, query id 90345 localhost root init
SHOW /*!50000 ENGINE*/ INNODB STATUS
---TRANSACTION 1093821554, not started
MySQL thread id 27893, OS thread handle 0x7f6c0c886700, query id 90144 127.0.0.1 cactiuser cleaning up
---TRANSACTION 1093821583, not started
MySQL thread id 27888, OS thread handle 0x7f6c0c8c7700, query id 90175 127.0.0.1 cactiuser cleaning up
---TRANSACTION 1093811214, not started
MySQL thread id 27887, OS thread handle 0x7f6c0c98a700, query id 80071 127.0.0.1 cactiuser cleaning up
---TRANSACTION 1093820819, not started
MySQL thread id 27886, OS thread handle 0x7f6c0c949700, query id 89403 127.0.0.1 cactiuser cleaning up
---TRANSACTION 1093811160, not started
MySQL thread id 27885, OS thread handle 0x7f6c0c908700, query id 80015 127.0.0.1 cactiuser cleaning up
--------
FILE I/O
--------
I/O thread 0 state: waiting for completed aio requests (insert buffer thread)
I/O thread 1 state: waiting for completed aio requests (log thread)
I/O thread 2 state: waiting for completed aio requests (read thread)
I/O thread 3 state: waiting for completed aio requests (read thread)
I/O thread 4 state: waiting for completed aio requests (read thread)
I/O thread 5 state: waiting for completed aio requests (read thread)
I/O thread 6 state: waiting for completed aio requests (write thread)
I/O thread 7 state: waiting for completed aio requests (write thread)
I/O thread 8 state: waiting for completed aio requests (write thread)
I/O thread 9 state: waiting for completed aio requests (write thread)
Pending normal aio reads: 0 [0, 0, 0, 0] , aio writes: 0 [0, 0, 0, 0] ,
 ibuf aio reads: 0, log i/o's: 0, sync i/o's: 0
Pending flushes (fsync) log: 0; buffer pool: 0
124669 OS file reads, 4457 OS file writes, 3498 OS fsyncs
0.00 reads/s, 0 avg bytes/read, 0.00 writes/s, 0.00 fsyncs/s
-------------------------------------
INSERT BUFFER AND ADAPTIVE HASH INDEX
-------------------------------------
Ibuf: size 1, free list len 63, seg size 65, 2 merges
merged operations:
 insert 48, delete mark 0, delete 0
discarded operations:
 insert 0, delete mark 0, delete 0
Hash table size 34679, node heap has 1 buffer(s)
0.00 hash searches/s, 0.00 non-hash searches/s
---
LOG
---
Log sequence number 53339891261
Log flushed up to   53339891261
Pages flushed up to 53339891261
Last checkpoint at  53339891261
10 pending log writes, 20 pending chkp writes
3395 log i/o's done, 0.00 log i/o's/second
----------------------
BUFFER POOL AND MEMORY
----------------------
Total memory allocated 17170432; in additional pool allocated 0
Dictionary memory allocated 318159
Buffer pool size   1024
Free buffers       755
Database pages     256
Old database pages 0
Modified db pages  0
Pending reads 0
Pending writes: LRU 0, flush list 0, single page 0
Pages made young 6, not young 751793
0.00 youngs/s, 0.00 non-youngs/s
Pages read 124617, created 40, written 1020
0.00 reads/s, 0.00 creates/s, 0.00 writes/s
No buffer pool page gets since the last printout
Pages read ahead 0.00/s, evicted without access 0.00/s, Random read ahead 0.00/s
LRU len: 256, unzip_LRU len: 0
I/O sum[0]:cur[0], unzip sum[0]:cur[0]
--------------
ROW OPERATIONS
--------------
0 queries inside InnoDB, 0 queries in queue
0 read views open inside InnoDB
Main thread process no. 1968, id 140101998331648, state: sleeping
Number of rows inserted 3089, updated 220, deleted 212, read 2099881
0.00 inserts/s, 0.00 updates/s, 0.00 deletes/s, 0.00 reads/s
----------------------------
END OF INNODB MONITOR OUTPUT`
	stat := make(map[string]float64)

	parseInnodbStatus(stub, stat)
	// Innodb Semaphores
	assert.EqualValues(t, stat["spin_waits"], 947)
	assert.EqualValues(t, stat["spin_rounds"], 9442)
	assert.EqualValues(t, stat["os_waits"], 222)
	assert.EqualValues(t, stat["innodb_sem_wait"], 0)         // empty
	assert.EqualValues(t, stat["innodb_sem_wait_time_ms"], 0) // empty
	// Innodb Transactions
	assert.EqualValues(t, stat["innodb_transactions"], 1093821584)
	assert.EqualValues(t, stat["unpurged_txns"], 6021)
	assert.EqualValues(t, stat["history_list"], 649)
	assert.EqualValues(t, stat["current_transactions"], 6)
	assert.EqualValues(t, stat["active_transactions"], 0)
	assert.EqualValues(t, stat["innodb_lock_wait_secs"], 0) // empty
	assert.EqualValues(t, stat["read_views"], 0)            // empty
	assert.EqualValues(t, stat["innodb_tables_in_use"], 0)  // empty
	assert.EqualValues(t, stat["innodb_locked_tables"], 0)  // empty
	assert.EqualValues(t, stat["locked_transactions"], 0)   // empty
	assert.EqualValues(t, stat["innodb_lock_structs"], 0)   // empty
	// File I/O
	assert.EqualValues(t, stat["pending_normal_aio_reads"], 0)
	assert.EqualValues(t, stat["pending_normal_aio_writes"], 0)
	assert.EqualValues(t, stat["pending_ibuf_aio_reads"], 0)
	assert.EqualValues(t, stat["pending_aio_log_ios"], 0)
	assert.EqualValues(t, stat["pending_aio_sync_ios"], 0)
	assert.EqualValues(t, stat["pending_log_flushes"], 0)
	assert.EqualValues(t, stat["pending_buf_pool_flushes"], 0)
	// Insert Buffer and Adaptive Hash Index
	assert.EqualValues(t, stat["ibuf_used_cells"], 1)
	assert.EqualValues(t, stat["ibuf_free_cells"], 63)
	assert.EqualValues(t, stat["ibuf_cell_count"], 65)
	assert.EqualValues(t, stat["ibuf_inserts"], 48)
	assert.EqualValues(t, stat["ibuf_merges"], 2)
	assert.EqualValues(t, stat["ibuf_merged"], 48)
	assert.EqualValues(t, stat["hash_index_cells_total"], 34679)
	assert.EqualValues(t, stat["hash_index_cells_used"], 0) // empty
	// Log
	assert.EqualValues(t, stat["log_writes"], 3395)
	assert.EqualValues(t, stat["pending_log_writes"], 10)
	assert.EqualValues(t, stat["pending_chkp_writes"], 20)
	assert.EqualValues(t, stat["log_bytes_written"], 53339891261)
	assert.EqualValues(t, stat["log_bytes_flushed"], 53339891261)
	assert.EqualValues(t, stat["last_checkpoint"], 53339891261)
	assert.EqualValues(t, stat["log_pending_log_flushes"], 0)
	// Buffer Pool and Memory
	assert.EqualValues(t, stat["total_mem_alloc"], 17170432)
	assert.EqualValues(t, stat["additional_pool_alloc"], 0)
	assert.EqualValues(t, stat["adaptive_hash_memory"], 0)     // empty
	assert.EqualValues(t, stat["page_hash_memory"], 0)         // empty
	assert.EqualValues(t, stat["dictionary_cache_memory"], 0)  // empty
	assert.EqualValues(t, stat["file_system_memory"], 0)       // empty
	assert.EqualValues(t, stat["lock_system_memory"], 0)       // empty
	assert.EqualValues(t, stat["recovery_system_memory"], 0)   // empty
	assert.EqualValues(t, stat["thread_hash_memory"], 0)       // empty
	assert.EqualValues(t, stat["innodb_io_pattern_memory"], 0) // empty
	// etc
	assert.EqualValues(t, stat["unflushed_log"], 0)
	assert.EqualValues(t, stat["uncheckpointed_bytes"], 0)

}

func TestParseProcStat57(t *testing.T) {
	stub := `
=====================================
2016-02-22 19:08:31 0x700000eda000 INNODB MONITOR OUTPUT
=====================================
Per second averages calculated from the last 4 seconds
-----------------
BACKGROUND THREAD
-----------------
srv_master_thread loops: 1 srv_active, 0 srv_shutdown, 2 srv_idle
srv_master_thread log flush and writes: 3
----------
SEMAPHORES
----------
OS WAIT ARRAY INFO: reservation count 63
OS WAIT ARRAY INFO: signal count 111
RW-shared spins 0, rounds 85, OS waits 22
RW-excl spins 0, rounds 4705, OS waits 17
RW-sx spins 70, rounds 70, OS waits 70
Spin rounds per wait: 85.00 RW-shared, 4705.00 RW-excl, 0.00 RW-sx
------------
TRANSACTIONS
------------
Trx id counter 49154
Purge done for trx's n:o < 44675 undo n:o < 0 state: running but idle
History list length 775
LIST OF TRANSACTIONS FOR EACH SESSION:
---TRANSACTION 281479529875248, not started
0 lock struct(s), heap size 1136, 0 row lock(s)
--------
FILE I/O
--------
I/O thread 0 state: waiting for i/o request (insert buffer thread)
I/O thread 1 state: waiting for i/o request (log thread)
I/O thread 2 state: waiting for i/o request (read thread)
I/O thread 3 state: waiting for i/o request (read thread)
I/O thread 4 state: waiting for i/o request (read thread)
I/O thread 5 state: waiting for i/o request (read thread)
I/O thread 6 state: waiting for i/o request (write thread)
I/O thread 7 state: waiting for i/o request (write thread)
I/O thread 8 state: waiting for i/o request (write thread)
I/O thread 9 state: waiting for i/o request (write thread)
Pending normal aio reads: [0, 0, 0, 0] , aio writes: [0, 0, 0, 0] ,
 ibuf aio reads:, log i/o's:, sync i/o's:
Pending flushes (fsync) log: 0; buffer pool: 0
516 OS file reads, 55 OS file writes, 9 OS fsyncs
128.97 reads/s, 20393 avg bytes/read, 13.75 writes/s, 2.25 fsyncs/s
-------------------------------------
INSERT BUFFER AND ADAPTIVE HASH INDEX
-------------------------------------
Ibuf: size 1, free list len 0, seg size 2, 0 merges
merged operations:
 insert 0, delete mark 0, delete 0
discarded operations:
 insert 0, delete mark 0, delete 0
Hash table size 276671, node heap has 2 buffer(s)
Hash table size 276671, node heap has 0 buffer(s)
Hash table size 276671, node heap has 0 buffer(s)
Hash table size 276671, node heap has 0 buffer(s)
Hash table size 276671, node heap has 1 buffer(s)
Hash table size 276671, node heap has 1 buffer(s)
Hash table size 276671, node heap has 0 buffer(s)
Hash table size 276671, node heap has 4 buffer(s)
276.93 hash searches/s, 835.29 non-hash searches/s
---
LOG
---
Log sequence number 379575319
Log flushed up to   379575319
Pages flushed up to 379575319
Last checkpoint at  379575310
10 pending log flushes, 20 pending chkp writes
12 log i/o's done, 3.00 log i/o's/second
----------------------
BUFFER POOL AND MEMORY
----------------------
Total large memory allocated 1099431936
Dictionary memory allocated 312184
Buffer pool size   65528
Free buffers       64999
Database pages     521
Old database pages 0
Modified db pages  0
Pending reads 0
Pending writes: LRU 0, flush list 0, single page 0
Pages made young 0, not young 0
0.00 youngs/s, 0.00 non-youngs/s
Pages read 487, created 34, written 36
121.72 reads/s, 8.50 creates/s, 9.00 writes/s
Buffer pool hit rate 974 / 1000, young-making rate 0 / 1000 not 0 / 1000
Pages read ahead 0.00/s, evicted without access 0.00/s, Random read ahead 0.00/s
LRU len: 521, unzip_LRU len: 0
I/O sum[0]:cur[0], unzip sum[0]:cur[0]
----------------------
INDIVIDUAL BUFFER POOL INFO
----------------------
---BUFFER POOL 0
Buffer pool size   16382
Free buffers       16228
Database pages     152
Old database pages 0
Modified db pages  0
Pending reads 0
Pending writes: LRU 0, flush list 0, single page 0
Pages made young 0, not young 0
0.00 youngs/s, 0.00 non-youngs/s
Pages read 152, created 0, written 2
37.99 reads/s, 0.00 creates/s, 0.50 writes/s
Buffer pool hit rate 976 / 1000, young-making rate 0 / 1000 not 0 / 1000
Pages read ahead 0.00/s, evicted without access 0.00/s, Random read ahead 0.00/s
LRU len: 152, unzip_LRU len: 0
I/O sum[0]:cur[0], unzip sum[0]:cur[0]
---BUFFER POOL 1
Buffer pool size   16382
Free buffers       16244
Database pages     136
Old database pages 0
Modified db pages  0
Pending reads 0
Pending writes: LRU 0, flush list 0, single page 0
Pages made young 0, not young 0
0.00 youngs/s, 0.00 non-youngs/s
Pages read 136, created 0, written 0
33.99 reads/s, 0.00 creates/s, 0.00 writes/s
Buffer pool hit rate 978 / 1000, young-making rate 0 / 1000 not 0 / 1000
Pages read ahead 0.00/s, evicted without access 0.00/s, Random read ahead 0.00/s
LRU len: 136, unzip_LRU len: 0
I/O sum[0]:cur[0], unzip sum[0]:cur[0]
---BUFFER POOL 2
Buffer pool size   16382
Free buffers       16313
Database pages     67
Old database pages 0
Modified db pages  0
Pending reads 0
Pending writes: LRU 0, flush list 0, single page 0
Pages made young 0, not young 0
0.00 youngs/s, 0.00 non-youngs/s
Pages read 67, created 0, written 0
16.75 reads/s, 0.00 creates/s, 0.00 writes/s
Buffer pool hit rate 975 / 1000, young-making rate 0 / 1000 not 0 / 1000
Pages read ahead 0.00/s, evicted without access 0.00/s, Random read ahead 0.00/s
LRU len: 67, unzip_LRU len: 0
I/O sum[0]:cur[0], unzip sum[0]:cur[0]
---BUFFER POOL 3
Buffer pool size   16382
Free buffers       16214
Database pages     166
Old database pages 0
Modified db pages  0
Pending reads 0
Pending writes: LRU 0, flush list 0, single page 0
Pages made young 0, not young 0
0.00 youngs/s, 0.00 non-youngs/s
Pages read 132, created 34, written 34
32.99 reads/s, 8.50 creates/s, 8.50 writes/s
Buffer pool hit rate 963 / 1000, young-making rate 0 / 1000 not 0 / 1000
Pages read ahead 0.00/s, evicted without access 0.00/s, Random read ahead 0.00/s
LRU len: 166, unzip_LRU len: 0
I/O sum[0]:cur[0], unzip sum[0]:cur[0]
--------------
ROW OPERATIONS
--------------
0 queries inside InnoDB, 0 queries in queue
0 read views open inside InnoDB
Process ID=28837, Main thread ID=123145312497664, state: sleeping
Number of rows inserted 0, updated 0, deleted 0, read 8
0.00 inserts/s, 0.00 updates/s, 0.00 deletes/s, 2.00 reads/s
----------------------------
END OF INNODB MONITOR OUTPUT
============================
`
	stat := make(map[string]float64)
	parseInnodbStatus(stub, stat)
	// Innodb Semaphores
	assert.EqualValues(t, stat["spin_waits"], 70)
	assert.EqualValues(t, stat["spin_rounds"], 0) // empty
	assert.EqualValues(t, stat["os_waits"], 109)
	assert.EqualValues(t, stat["innodb_sem_wait"], 0)         // empty
	assert.EqualValues(t, stat["innodb_sem_wait_time_ms"], 0) // empty
	// Innodb Transactions
	assert.EqualValues(t, stat["innodb_transactions"], 49154) // empty
	assert.EqualValues(t, stat["unpurged_txns"], 4479)
	assert.EqualValues(t, stat["history_list"], 775)
	assert.EqualValues(t, stat["current_transactions"], 1)
	assert.EqualValues(t, stat["active_transactions"], 0)
	assert.EqualValues(t, stat["innodb_lock_wait_secs"], 0) // empty
	assert.EqualValues(t, stat["read_views"], 0)
	assert.EqualValues(t, stat["innodb_tables_in_use"], 0) // empty
	assert.EqualValues(t, stat["innodb_locked_tables"], 0) // empty
	assert.EqualValues(t, stat["locked_transactions"], 0)  // empty
	assert.EqualValues(t, stat["innodb_lock_structs"], 0)  // empty
	// File I/O
	assert.EqualValues(t, stat["pending_normal_aio_reads"], 0)
	assert.EqualValues(t, stat["pending_normal_aio_writes"], 0)
	assert.EqualValues(t, stat["pending_ibuf_aio_reads"], 0)
	assert.EqualValues(t, stat["pending_aio_log_ios"], 0)
	assert.EqualValues(t, stat["pending_aio_sync_ios"], 0)
	assert.EqualValues(t, stat["pending_log_flushes"], 0)
	assert.EqualValues(t, stat["pending_buf_pool_flushes"], 0)
	// Insert Buffer and Adaptive Hash Index
	assert.EqualValues(t, stat["ibuf_used_cells"], 1)
	assert.EqualValues(t, stat["ibuf_free_cells"], 0)
	assert.EqualValues(t, stat["ibuf_cell_count"], 2)
	assert.EqualValues(t, stat["ibuf_inserts"], 0)
	assert.EqualValues(t, stat["ibuf_merges"], 0)
	assert.EqualValues(t, stat["ibuf_merged"], 0)
	assert.EqualValues(t, stat["hash_index_cells_total"], 276671)
	assert.EqualValues(t, stat["hash_index_cells_used"], 0)
	// Log
	assert.EqualValues(t, stat["log_writes"], 12)
	assert.EqualValues(t, stat["pending_log_writes"], 0)
	assert.EqualValues(t, stat["pending_chkp_writes"], 20)
	assert.EqualValues(t, stat["log_bytes_written"], 379575319)
	assert.EqualValues(t, stat["log_bytes_flushed"], 379575319)
	assert.EqualValues(t, stat["last_checkpoint"], 379575310)
	assert.EqualValues(t, stat["log_pending_log_flushes"], 10)
	// Buffer Pool and Memory
	assert.EqualValues(t, stat["total_mem_alloc"], 1099431936)
	assert.EqualValues(t, stat["additional_pool_alloc"], 0)
	assert.EqualValues(t, stat["adaptive_hash_memory"], 0)     // empty
	assert.EqualValues(t, stat["page_hash_memory"], 0)         // empty
	assert.EqualValues(t, stat["dictionary_cache_memory"], 0)  // empty
	assert.EqualValues(t, stat["file_system_memory"], 0)       // empty
	assert.EqualValues(t, stat["lock_system_memory"], 0)       // empty
	assert.EqualValues(t, stat["recovery_system_memory"], 0)   // empty
	assert.EqualValues(t, stat["thread_hash_memory"], 0)       // empty
	assert.EqualValues(t, stat["innodb_io_pattern_memory"], 0) // empty
	// etc
	assert.EqualValues(t, stat["unflushed_log"], 0)
	assert.EqualValues(t, stat["uncheckpointed_bytes"], 9)

}

func TestParseLockedTransactions(t *testing.T) {

	stub := `=====================================
170829 11:50:33 INNODB MONITOR OUTPUT
=====================================
Per second averages calculated from the last 18 seconds
-----------------
BACKGROUND THREAD
-----------------
srv_master_thread loops: 26 1_second, 26 sleeps, 2 10_second, 9 background, 9 flush
srv_master_thread log flush and writes: 28
----------
SEMAPHORES
----------
OS WAIT ARRAY INFO: reservation count 12, signal count 11
Mutex spin waits 6, rounds 180, OS waits 6
RW-shared spins 6, rounds 180, OS waits 6
RW-excl spins 0, rounds 0, OS waits 0
Spin rounds per wait: 30.00 mutex, 30.00 RW-shared, 0.00 RW-excl
--------
FILE I/O
--------
I/O thread 0 state: waiting for completed aio requests (insert buffer thread)
I/O thread 1 state: waiting for completed aio requests (log thread)
I/O thread 2 state: waiting for completed aio requests (read thread)
I/O thread 3 state: waiting for completed aio requests (read thread)
I/O thread 4 state: waiting for completed aio requests (read thread)
I/O thread 5 state: waiting for completed aio requests (read thread)
I/O thread 6 state: waiting for completed aio requests (write thread)
I/O thread 7 state: waiting for completed aio requests (write thread)
I/O thread 8 state: waiting for completed aio requests (write thread)
I/O thread 9 state: waiting for completed aio requests (write thread)
Pending normal aio reads: 0 [0, 0, 0, 0] , aio writes: 0 [0, 0, 0, 0] ,
 ibuf aio reads: 0, log i/o's: 0, sync i/o's: 0
Pending flushes (fsync) log: 0; buffer pool: 0
310 OS file reads, 174 OS file writes, 22 OS fsyncs
0.00 reads/s, 0 avg bytes/read, 0.00 writes/s, 0.00 fsyncs/s
-------------------------------------
INSERT BUFFER AND ADAPTIVE HASH INDEX
-------------------------------------
Ibuf: size 1, free list len 0, seg size 2, 0 merges
merged operations:
 insert 0, delete mark 0, delete 0
discarded operations:
 insert 0, delete mark 0, delete 0
Hash table size 276671, node heap has 1 buffer(s)
0.00 hash searches/s, 0.00 non-hash searches/s
---
LOG
---
Log sequence number 1602283
Log flushed up to   1602283
Last checkpoint at  1602283
Max checkpoint age    7782360
Checkpoint age target 7539162
Modified age          0
Checkpoint age        0
10 pending log writes, 20 pending chkp writes
40 log i/o's done, 0.00 log i/o's/second
----------------------
BUFFER POOL AND MEMORY
----------------------
Total memory allocated 137756672; in additional pool allocated 0
Total memory allocated by read views 88
Internal hash tables (constant factor + variable factor)
    Adaptive hash index 2233968 	(2213368 + 20600)
    Page hash           139112 (buffer pool 0 only)
    Dictionary cache    597886 	(554768 + 43118)
    File system         83536 	(82672 + 864)
    Lock system         334000 	(332872 + 1128)
    Recovery system     0 	(0 + 0)
Dictionary memory allocated 43118
Buffer pool size        8191
Buffer pool size, bytes 134201344
Free buffers            8039
Database pages          151
Old database pages      0
Modified db pages       0
Pending reads 0
Pending writes: LRU 0, flush list 0, single page 0
Pages made young 0, not young 0
0.00 youngs/s, 0.00 non-youngs/s
Pages read 147, created 4, written 156
0.00 reads/s, 0.00 creates/s, 0.00 writes/s
No buffer pool page gets since the last printout
Pages read ahead 0.00/s, evicted without access 0.00/s, Random read ahead 0.00/s
LRU len: 151, unzip_LRU len: 0
I/O sum[0]:cur[0], unzip sum[0]:cur[0]
--------------
ROW OPERATIONS
--------------
0 queries inside InnoDB, 0 queries in queue
1 read views open inside InnoDB
2 transactions active inside InnoDB
2 out of 1000 descriptors used
---OLDEST VIEW---
Normal read view
Read view low limit trx n:o 505
Read view up limit trx id 505
Read view low limit trx id 505
Read view individually stored trx ids:
-----------------
Main thread process no. 458, id 139631366485760, state: waiting for server activity
Number of rows inserted 2, updated 0, deleted 1, read 2
0.00 inserts/s, 0.00 updates/s, 0.00 deletes/s, 0.00 reads/s
------------
TRANSACTIONS
------------
Trx id counter 507
Purge done for trx's n:o < 505 undo n:o < 0
History list length 1
LIST OF TRANSACTIONS FOR EACH SESSION:
---TRANSACTION 0, not started
MySQL thread id 8, OS thread handle 0x7efe7cb12700, query id 52 localhost root
SHOW ENGINE INNODB STATUS
---TRANSACTION 506, ACTIVE 804 sec starting index read
mysql tables in use 1, locked 1
LOCK WAIT 2 lock struct(s), heap size 376, 1 row lock(s)
MySQL thread id 3, OS thread handle 0x7efe7cb5b700, query id 47 localhost root statistics
SELECT * FROM test WHERE id = 1 LOCK IN SHARE MODE
------- TRX HAS BEEN WAITING 22 SEC FOR THIS LOCK TO BE GRANTED:
RECORD LOCKS space id 0 page no 307 n bits 72 index ` + "`PRIMARY` of table `test`.`test`" + ` trx id 506 lock mode S locks rec but not gap waiting
------------------
---TRANSACTION 505, ACTIVE 815 sec
2 lock struct(s), heap size 376, 1 row lock(s), undo log entries 1
MySQL thread id 2, OS thread handle 0x7efe7cba4700, query id 35 localhost root
----------------------------
END OF INNODB MONITOR OUTPUT
============================`
	stat := make(map[string]float64)
	parseInnodbStatus(stub, stat)
	// Innodb Semaphores
	assert.EqualValues(t, stat["spin_waits"], 12)
	assert.EqualValues(t, stat["spin_rounds"], 180)
	assert.EqualValues(t, stat["os_waits"], 12)
	assert.EqualValues(t, stat["innodb_sem_wait"], 0)         // empty
	assert.EqualValues(t, stat["innodb_sem_wait_time_ms"], 0) // empty
	// Innodb Transactions
	assert.EqualValues(t, stat["innodb_transactions"], 507)
	assert.EqualValues(t, stat["unpurged_txns"], 2)
	assert.EqualValues(t, stat["history_list"], 1)
	assert.EqualValues(t, stat["current_transactions"], 3)
	assert.EqualValues(t, stat["active_transactions"], 2)
	assert.EqualValues(t, stat["innodb_lock_wait_secs"], 22)
	assert.EqualValues(t, stat["read_views"], 1)
	assert.EqualValues(t, stat["innodb_tables_in_use"], 1)
	assert.EqualValues(t, stat["innodb_locked_tables"], 1)
	assert.EqualValues(t, stat["locked_transactions"], 1)
	assert.EqualValues(t, stat["innodb_lock_structs"], 4)
	// File I/O
	assert.EqualValues(t, stat["pending_normal_aio_reads"], 0)
	assert.EqualValues(t, stat["pending_normal_aio_writes"], 0)
	assert.EqualValues(t, stat["pending_ibuf_aio_reads"], 0)
	assert.EqualValues(t, stat["pending_aio_log_ios"], 0)
	assert.EqualValues(t, stat["pending_aio_sync_ios"], 0)
	assert.EqualValues(t, stat["pending_log_flushes"], 0)
	assert.EqualValues(t, stat["pending_buf_pool_flushes"], 0)
	// Insert Buffer and Adaptive Hash Index
	assert.EqualValues(t, stat["ibuf_used_cells"], 1)
	assert.EqualValues(t, stat["ibuf_free_cells"], 0)
	assert.EqualValues(t, stat["ibuf_cell_count"], 2)
	assert.EqualValues(t, stat["ibuf_inserts"], 0)
	assert.EqualValues(t, stat["ibuf_merges"], 0)
	assert.EqualValues(t, stat["ibuf_merged"], 0)
	assert.EqualValues(t, stat["hash_index_cells_total"], 276671)
	assert.EqualValues(t, stat["hash_index_cells_used"], 0)
	// Log
	assert.EqualValues(t, stat["log_writes"], 40)
	assert.EqualValues(t, stat["pending_log_writes"], 10)
	assert.EqualValues(t, stat["pending_chkp_writes"], 20)
	assert.EqualValues(t, stat["log_bytes_written"], 1602283)
	assert.EqualValues(t, stat["log_bytes_flushed"], 1602283)
	assert.EqualValues(t, stat["last_checkpoint"], 1602283)
	assert.EqualValues(t, stat["log_pending_log_flushes"], 0)
	// Buffer Pool and Memory
	assert.EqualValues(t, stat["total_mem_alloc"], 137756672)
	assert.EqualValues(t, stat["additional_pool_alloc"], 0)
	assert.EqualValues(t, stat["adaptive_hash_memory"], 2233968)
	assert.EqualValues(t, stat["page_hash_memory"], 139112)
	assert.EqualValues(t, stat["dictionary_cache_memory"], 597886)
	assert.EqualValues(t, stat["file_system_memory"], 83536)
	assert.EqualValues(t, stat["lock_system_memory"], 334000)
	assert.EqualValues(t, stat["recovery_system_memory"], 0)   // empty
	assert.EqualValues(t, stat["thread_hash_memory"], 0)       // empty
	assert.EqualValues(t, stat["innodb_io_pattern_memory"], 0) // empty
	// etc
	assert.EqualValues(t, stat["unflushed_log"], 0)
	assert.EqualValues(t, stat["uncheckpointed_bytes"], 0)

}

func TestParseProcesslist1(t *testing.T) {
	stat := make(map[string]float64)
	pattern := []string{"NULL"}

	for _, val := range pattern {
		parseProcesslist(val, stat)
	}
	assert.EqualValues(t, 0, stat["State_closing_tables"])
	assert.EqualValues(t, 0, stat["State_copying_to_tmp_table"])
	assert.EqualValues(t, 0, stat["State_end"])
	assert.EqualValues(t, 0, stat["State_freeing_items"])
	assert.EqualValues(t, 0, stat["State_init"])
	assert.EqualValues(t, 0, stat["State_locked"])
	assert.EqualValues(t, 0, stat["State_login"])
	assert.EqualValues(t, 0, stat["State_preparing"])
	assert.EqualValues(t, 0, stat["State_reading_from_net"])
	assert.EqualValues(t, 0, stat["State_sending_data"])
	assert.EqualValues(t, 0, stat["State_sorting_result"])
	assert.EqualValues(t, 0, stat["State_statistics"])
	assert.EqualValues(t, 0, stat["State_updating"])
	assert.EqualValues(t, 0, stat["State_writing_to_net"])
	assert.EqualValues(t, 0, stat["State_none"])
	assert.EqualValues(t, 1, stat["State_other"])
}

func TestParseProcesslist2(t *testing.T) {
	stat := make(map[string]float64)

	// https://dev.mysql.com/doc/refman/5.6/en/general-thread-states.html
	pattern := []string{
		"",
		"After create",
		"altering table",
		"Analyzing",
		"checking permissions",
		"Checking table",
		"cleaning up",
		"closing tables",
		"committing alter table to storage engine",
		"converting HEAP to MyISAM",
		"MEMORY",
		"MyISAM",
		"copy to tmp table",
		"Copying to group table",
		"GROUP BY",
		"Copying to tmp table",
		"Copying to tmp table on disk",
		"Creating index",
		"Creating sort index",
		"creating table",
		"Creating tmp table",
		"deleting from main table",
		"deleting from reference tables",
		"discard_or_import_tablespace",
		"end",
		"executing",
		"Execution of init_command",
		"freeing items",
		"FULLTEXT initialization",
		"init",
		"Killed",
		"logging slow query",
		"login",
		"manage keys",
		"NULL",
		"Opening tables",
		"Opening table",
		"optimizing",
		"preparing",
		"preparing for alter table",
		"Purging old relay logs",
		"query end",
		"Reading from net",
		"Removing duplicates",
		"removing tmp table",
		"rename",
		"rename result table",
		"Reopen tables",
		"Repair by sorting",
		"Repair done",
		"Repair with keycache",
		"Rolling back",
		"Saving state",
		"Searching rows for update",
		"Sending data",
		"setup",
		"Sorting for group",
		"Sorting for order",
		"Sorting index",
		"Sorting result",
		"statistics",
		"System lock",
		"update",
		"Updating",
		"updating main table",
		"updating reference tables",
		"User lock",
		"User sleep",
		"Waiting for commit lock",
		"Waiting for global read lock",
		"Waiting for tables",
		"Waiting for table flush",
		"Waiting for lock_type lock",
		"Waiting for table level lock",
		"Waiting for event metadata lock",
		"Waiting for global read lock",
		"Waiting for schema metadat lock",
		"Waiting for stored function metadata  lock",
		"Waiting for stored procedure metadata lock",
		"Waiting for table metadata lock",
		"Waiting for trigger metadata lock",
		"Waiting on cond",
		"Writing to net",
		"Table lock",
	}

	for _, val := range pattern {
		parseProcesslist(val, stat)
	}
	assert.EqualValues(t, 1, stat["State_closing_tables"])
	assert.EqualValues(t, 1, stat["State_copying_to_tmp_table"])
	assert.EqualValues(t, 1, stat["State_end"])
	assert.EqualValues(t, 1, stat["State_freeing_items"])
	assert.EqualValues(t, 1, stat["State_init"])
	assert.EqualValues(t, 12, stat["State_locked"])
	assert.EqualValues(t, 1, stat["State_login"])
	assert.EqualValues(t, 1, stat["State_preparing"])
	assert.EqualValues(t, 1, stat["State_reading_from_net"])
	assert.EqualValues(t, 1, stat["State_sending_data"])
	assert.EqualValues(t, 1, stat["State_sorting_result"])
	assert.EqualValues(t, 1, stat["State_statistics"])
	assert.EqualValues(t, 1, stat["State_updating"])
	assert.EqualValues(t, 1, stat["State_writing_to_net"])
	assert.EqualValues(t, 1, stat["State_none"])
	assert.EqualValues(t, 58, stat["State_other"])
}

type TestCaseAio struct {
	stub   string
	reads  int
	writes int
}

func TestParseAio(t *testing.T) {
	pattern := []TestCaseAio{
		{"Pending normal aio reads: [1, 3, 5, 7] , aio writes: [3, 5, 7, 9] ,", 16, 24},
		{"Pending normal aio reads: [1, 3, 5, 7] ", 16, 0},
		{"Pending normal aio reads: 10 [4, 6] , aio writes: 20 [2, 4, 6, 8] ,", 10, 20},
		{"Pending normal aio reads: 10 [4, 6] ", 10, 0},
		{"Pending normal aio reads: 10, aio writes: 20,", 10, 20},
		{"Pending normal aio reads: 10", 10, 0},
		{"Pending normal aio reads:, aio writes: [1, 3, 5, 7],", 0, 16},
		{"Pending normal aio reads:, aio writes:,", 0, 0},
	}

	for _, tt := range pattern {
		stat := make(map[string]float64)
		parseInnodbStatus(tt.stub, stat)
		assert.EqualValues(t, stat["pending_normal_aio_reads"], tt.reads)
		assert.EqualValues(t, stat["pending_normal_aio_writes"], tt.writes)
	}
}

func TestMetricNamesShouldUniqueAndConst(t *testing.T) {
	m := MySQLPlugin{
		DisableInnoDB:  false,
		EnableExtended: true,
	}
	defs := m.GraphDefinition()
	keys := make(map[string]string) // metricName: graphDefName
	for name, g := range defs {
		for _, v := range g.Metrics {
			if v.Name == "Threads_connected" {
				if name != "connections" && name != "threads" {
					t.Errorf(`%q are duplicated in "connections", "threads" and %q`, v.Name, name)
				}
				continue
			}
			if v.Name == "Qcache_hits" {
				if name != "cmd" && name != "query_cache" {
					t.Errorf(`%q are duplicated in "cmd", "query_cache" and %q`, v.Name, name)
				}
				continue
			}

			if strings.ContainsAny(v.Name, "#*") {
				t.Errorf("%q should not contains wildcards", v.Name)
			}
			if s, ok := keys[v.Name]; ok {
				t.Errorf("%q are defined in both %q and %q", v.Name, s, name)
			}
			keys[v.Name] = name
		}
	}
}
