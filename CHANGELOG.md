# Change history

## Feb 17 2015 : v1.4.0

  This is a major release, and makes using the client much easier to develop applications.

  * **New Features**

    * Added Marshalling Support for Put and Get operations. Refer to [Marshalling Test](client_object_test.go) to see how to take advantage.
    Same functionality for other APIs will follow soon.
    Example:
    ```go
    type SomeStruct struct {
      A    int            `as:"a"`  // alias the field to myself
      Self *SomeStruct    `as:"-"`  // will not persist the field
    }

    type OtherStruct struct {
      i interface{}
      OtherObject *OtherStruct
    }

    obj := &OtherStruct {
      i: 15,
      OtherObject: OtherStruct {A: 18},
    }

    key, _ := as.NewKey("ns", "set", value)
    err := client.PutObject(nil, key, obj)
    // handle error here

    rObj := &OtherStruct{}
    err = client.GetObject(nil, key, rObj)
    ```

    * Added `Recordset.Results()`. Consumers of a recordset do not have to implement a select anymore. Instead of:
    ```go
    recordset, err := client.ScanAll(...)
    L:
    for {
      select {
      case r := <-recordset.Record:
        if r == nil {
          break L
        }
        // process record here
      case e := <-recordset.Errors:
        // handle error here
      }
    }
    ```

    one should only range on `recordset.Results()`:

    ```go
    recordset, err := client.ScanAll(...)
    for res := range recordset.Results() {
      if res.Err != nil {
        // handle error here
      } else {
        // process record here
        fmt.Println(res.Record.Bins)
      }
    }
    ```

    Use of the old pattern is discouraged and deprecated, and direct access to recordset.Records and recordset.Errors will be removed in a future release.
      
  * **Improvements**

    * Custom Types are now allowed as bin values.

## Jan 26 2015 : v1.3.1

  * **Improvements**

    * Removed dependency on `unsafe` package.

## Jan 20 2015 : v1.3.0

  * **Breaking Changes**

    * Removed `Record.Duplicates` and `GenerationPolicy/DUPLICATE`

  * **New Features**

    * Added Security Features: Please consult [Security Docs](https://www.aerospike.com/docs/guide/security.html) on Aerospike website.
      
      * `ClientPolicy.User`, `ClientPolicy.Password`
      * `Client.CreateUser()`, `Client.DropUser()`, `Client.ChangePassword()`
      * `Client.GrantRoles()`, `Client.RevokeRoles()`, `Client.ReplaceRoles()`
      * `Client.QueryUser()`, `Client.QueryUsers`

    * Added `Client.QueryNode()`

    * Added `ClientPolicy.TendInterval`

  * **Improvements**

    * Cleaned up Scan/Query/Recordset concurrent code

  * **Fixes**

      * Fixed a bug in `tools/cli/cli.go`.

      * Fixed a bug when `GetHeaderOp()` would always translate into `GetOp()`

## Dec 29 2014: v1.2.0

  * **New Features**

    * Added `NewKeyWithDigest()` method. You can now create keys with custom digests, or only using digests without
      knowing the original value. (Useful when you are getting back results with Query and Scan)

## Dec 22 2014

  * **New Features**

    * Added `ConsistencyLevel` to `BasePolicy`.

    * Added `CommitLevel` to `WritePolicy`.

    * Added `LargeList.Range` and `LargeList.RangeThenFilter` methods.

    * Added `LargeMap.Exists` method.

  * **Improvements**

    * We use a pooled XORShift RNG to produce random numbers in the client. It is FAST.

## Dec 19 2014

  * **Fixes**

    * `Record.Expiration` wasn't converted to TTL values on `Client.BatchGet`, `Client.Scan` and `Client.Query`.

## Dec 10 2014

  * **Fixes**:

    * Fixed issue when the size of key field would not be estimated correctly when WritePolicy.SendKey was set.

## Nov 27 2014

  Major Performance Enhancements. Minor new features and fixes.

  * **Improvements**

    * Go client is much faster and more memory efficient now.
      In some workloads, it competes and wins against C and Java clients.

    * Complex objects are now de/serialized much faster.

  * **New Features**

    * Added Default Policies for Client object.
      Instead of creating a new policy when the passed policy is nil, default policies will be used.

## Nov 24 2014

  * **Fixes**:

    * Fixed issue when WritePolicy.SendKey = true was not respected in Touch() and Operate()

## Nov 22 2014

  Hotfix in unpacker. Update strongly recommended for everyone using Complex objects, LDTs and UDFs.

  * **Fixes**:

    * When Blob, ByteArray or String size has a bit sign set, unpacker reads it wrong.
        Note: This bug only affects unpacking of these objects. Packing was unaffected, and data in the database is valid.

## Nov 2 2014

  Minor, but very impoortant fix.

  * **Fixes**:

    * Node selection in partition map was flawed on first refresh.

  * **Incompatible changes**:

    * `Expiration` and `Generation` in `WritePolicy` are now `int32`
    * `TaskId` in `Statement` is now always set in the client, and is `int64`

  * **New Features**:

    * float32, float64 and bool are now supported in map and array types

## Oct 15 2014 (Beta 2)

  * **Hot fix**:

    * Fixed pack/unpack for uint64

## Aug 20 2014 (Beta 1)

  Major changes and improvements.

  * **New Features**:

    * Added client.Query()
    * Added client.ScanNode()/All()
    * Added client.Operate()
    * Added client.CreateIndex()
    * Added client.DropIndex()
    * Added client.RegisterUDF()
    * Added client.RegisterUDFFromFile()
    * Added client.Execute()
    * Added client.ExecuteUDF()
    * Added client.BatchGet()
    * Added client.BatchGetHeader()
    * Added client.BatchExists()
    * Added LDT implementation
    * Added `Node` and `Key` references to the Record

  * **Changes**:

    * Many minor and major bug fixes
    * Potentially breaking change: Reduced Undocumented API surface
    * Fixed a few places where error results were not checked
    * Breaking Change: Convert Key.namespace & Key.setName from pointer to string; affects Key API
    * Renamed all `this` receivers to appropriate names
    * Major performance improvements (~2X improvements in speed and memory consumption):
      * better memory management for commands; won't allocate if capacity is big enough
      * better hash management in key; avoids two redundant memory allocs
      * use a buffer pool to reduce GC load
      * fine-grained, customizable and deterministic buffer pool implementation for command

    * Optimizations for Key & Digest
      * changed digest implementation, removed an allocation
      * Added RIPEMD160 hash files from crypto to lib
      * pool hash objects

    * Various Benchmark tool improvements
      * now profileable using localhost:6060
      * minor bug fixes

## Jul 26 2014 (Alpha)

  * Initial Release.
