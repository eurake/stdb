Test tdengine 

// 创建超级表
CREATE TABLE thermometer (ts timestamp, degree double) TAGS(location binary(20), type int);

// 多功能表
// 复费率
CREATE TABLE thermometer (ts timestamp, 
    E1 float, E2 float, E3 float, E4 float, E5 float,  // 正有
    F1 float, F2 float, F3 float, F4 float, F5 float,  // 反有
    Z1 float, Z2 float, Z3 float, Z4 float, Z5 float,  // 正无
    X1 float, X2 float, X3 float, X4 float, X5 float,  // 反无
    UA float, UB float, UC float, (V)
    IA float, IB float, IC float, (A)
    QA float, QB float, QC float, // kW
    QTotal float,  // kW
    WA float, WB float, WC float, // kvar
    WTotal float,  // kvar
    SA float, SB float, SC float, // kVA
    STotal float,  // kVA
    GA float, GB float, GC float, // 功率因素
    GTotal float,  // 功率因素
    P float, // 电网频率
    OnOff bool
    ) TAGS(Num binary(20), type int, P1 float, P2 float, P3 float, P4 float, P5 float, );



// 多功能表
// DTSD DTSF 1620
CREATE TABLE thermometer (ts timestamp, 
    E1 float, E2 float, E3 float, E4 float, E5 float,  // 正有
    F1 float, F2 float, F3 float, F4 float, F5 float,  // 反有
    Z1 float, Z2 float, Z3 float, Z4 float, Z5 float,  // 正无
    X1 float, X2 float, X3 float, X4 float, X5 float,  // 反无
    UA float, UB float, UC float, (V)
    IA float, IB float, IC float, (A)
    QA float, QB float, QC float, // kW
    QTotal float,  // kW
    WA float, WB float, WC float, // kvar
    WTotal float,  // kvar
    SA float, SB float, SC float, // kVA
    STotal float,  // kVA
    GA float, GB float, GC float, // 功率因素
    GTotal float,  // 功率因素
    P float, // 电网频率
    OnOff bool
    ) TAGS(Num binary(20), type int);

// DTS1620
CREATE TABLE DTS1620 (ts timestamp, 
    E1 float, 
    UA float, UB float, UC float, (V)
    IA float, IB float, IC float, (A)
    QA float, QB float, QC float, // kW
    QTotal float,  // kW 总有功功率
    GA float, GB float, GC float, // 功率因素
    GTotal float,  // 功率因素
    P float, // 电网频率
    OnOff bool
    ) TAGS(Num binary(20), type int);


// Water
CREATE TABLE Water (ts timestamp, 
    E1 float, 
    OnOff bool
    ) TAGS(Num binary(20), type int);


// PayHistory
CREATE TABLE PayHistory (ts timestamp,
    S timestamp,
    E timestamp,
    S1 float,
    S2 float,
    S3 float,
    S4 float,
    S5 float,
    E1 float,
    E2 float,
    E3 float,
    E4 float,
    E5 float,
    U1 float,
    U2 float,
    U3 float,
    U4 float,
    U5 float,
    AferMoney float,
    BeforeMoney float,
    UsedMoney float
    M int
    ) TAGS(Num binary(20), type int);




// 为每个设备创建一个表
CREATE TABLE therm1 USING thermometer TAGS("beijing", 1);
CREATE TABLE therm2 USING thermometer TAGS("beijing", 2);
// 更新tag 
ALTER TABLE therm2 SET TAG location="shanghai";


// 插入数据
INSERT INTO therm1 VALUES("2020-01-01 00:00:00.000", 21);
INSERT INTO therm2 VALUES("2020-01-01 00:00:00.000", 23);

// 查询
SELECT COUNT(*), AVG(degree), MAX(degree), MIN(degree)
FROM thermometer
WHERE name<>'beijing' and ts>=now-1d
INTERVAL(10M)
GROUP BY location, type

// 其它需要用到的 当删除一个设备时用到
DROP TABLE therm1; 

// 

// 查询超级表下所有的子表数量
SELECT COUNT(TBNAME) FROM meters;
// 查询子表下所有标签
SELECT TBNAME, location FROM meters;

// 查询所有超级表下所有子表的最后一次数据更新
SELECT LAST_ROW(*) FROM st GROUP BY meters, tag1, tag2….

// 查询某个时间段的最后一条数据
SELECT LAST(*) FROM st WHERE ts > '2018-06-01 08:00:00.000' AND ts <= '2018-06-02 08:00:00.000' GROUP BY meters, tag1, tag2….

// 查询一个设备下某个时间段的数据
SELECT * FROM therm1 WHERE ts > '2018-06-01 08:00:00.000' AND ts <= '2018-06-02 08:00:00.000' ORDER BY ts DESC;
// 

由于插入数据不再要求按时间顺序,可以用以下方法实现删除某条记录或更新某条记录
步骤1: 查询某个表的数据, 缓存
步骤2: 修改或删除某条件缓存数据
步骤3: 删除这个表
步骤4: 创建这个表
步骤5: 缓存数据存储