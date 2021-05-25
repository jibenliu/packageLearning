package main

import (
	"archive/zip"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	srcDir        = "e:\\MSSQLBACKUP"
	outDir        = "e:\\MSSQLBACKUP"
	ZIPOutputPath = "e:\\MSSQLBACKUP"
)

type Table struct {
	s string
}

func main() {
	var server = "10.5.10.165\\sql2016"
	var port = 1433
	var user = "sa"
	var password = "95938"
	var database = "dotnet_zzsgl_yz_cs"


	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", server, port, database, user, password)
	fmt.Println(connString)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}

	defer conn.Close()

	// SELECT * FROM sysobjects WHERE TYPE='P' and name='sp_helptable'

	tableStructure := `
CREATE Procedure sp_helptable
(
    @table varchar(100)
)
-- exec sp_helptable tablename
-- 增加获取注释信息(感谢 袁罗)
AS 
Begin
declare @sql table(s varchar(1000), id int identity)
-- 创建语句
insert into  @sql(s) values ('create table [' + @table + '] (')

--获取注释
SELECT
A.name AS table_name,
B.name AS column_name,
C.value AS column_description
into #columnsproperties
FROM sys.tables A
INNER JOIN sys.columns B ON B.object_id = A.object_id
LEFT JOIN sys.extended_properties C ON C.major_id = B.object_id AND C.minor_id = B.column_id
WHERE A.name = @table

-- 获取列的列表，拼接语句
insert into @sql(s)
select 
    '  ['+a.column_name+'] ' + 
    data_type + coalesce('('+cast(character_maximum_length as varchar)+')','') + ' ' +
    case when exists ( 
        select id from syscolumns
        where object_name(id)=@table
        and name=a.column_name
        and columnproperty(id,name,'IsIdentity') = 1 
    ) then
        'IDENTITY(' + 
        cast(ident_seed(@table) as varchar) + ',' + 
        cast(ident_incr(@table) as varchar) + ')'
    else ''
    end + ' ' +
    ( case when IS_NULLABLE = 'NO' then 'NOT ' else '' end ) + 'NULL ' + 
    coalesce('DEFAULT '+COLUMN_DEFAULT,'') + case when isnull(convert(varchar,b.column_description),'')<>'' then  '/**'+isnull(convert(varchar,b.column_description),'')+'**/,'
else ',' end
 from INFORMATION_SCHEMA.COLUMNS  a left join #columnsproperties b  
 on convert(varchar,a.column_name)=convert(varchar,b.column_name)
 where a.table_name = @table
 order by ordinal_position
-- 主键
declare @pkname varchar(100)
select @pkname = constraint_name from INFORMATION_SCHEMA.TABLE_CONSTRAINTS
where table_name = @table and constraint_type='PRIMARY KEY'
if ( @pkname is not null ) begin
    insert into @sql(s) values('  PRIMARY KEY (')
    insert into @sql(s)
        select '   ['+COLUMN_NAME+'],' from INFORMATION_SCHEMA.KEY_COLUMN_USAGE
        where constraint_name = @pkname
        order by ordinal_position
    -- 去除尾部多余的字符
    update @sql set s=left(s,len(s)-1) where id=@@identity
    insert into @sql(s) values ('  )')
end
else begin
    -- 去除尾部多余的字符
    update @sql set s=left(s,len(s)-1) where id=@@identity
end
-- 继续拼接
insert into @sql(s) values( ')' )
-- 输出结果
select s from @sql order by id
END`
	stmt, err := conn.Prepare(tableStructure)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	insertProcedure := `
SET QUOTED_IDENTIFIER ON
SET ANSI_NULLS ON
GO
CREATE PROCEDURE [dbo].[spGenInsertSQL]
    (
      @tablename VARCHAR(256) ,
      @number BIGINT ,
      @whereClause NVARCHAR(MAX)
    )
AS
    BEGIN
        CREATE TABLE #temp ( [sql] NVARCHAR(MAX) )
    
        DECLARE @sql VARCHAR(MAX)
        DECLARE @sqlValues VARCHAR(MAX)
        SET @sql = ' ('
        SET @sqlValues = 'values (''+'
        SELECT  @sqlValues = @sqlValues + Cols + ' + '','' + ' ,
                @sql = @sql + '[' + name + '],'
        FROM    ( SELECT    CASE WHEN xtype IN ( 48, 52, 56, 59, 60, 62, 104,
                                                 106, 108, 122, 127 )
                                 THEN 'case when [' + name
                                      + '] is null then ''NULL'' else '
                                      + 'cast([' + name + '] as varchar)'
                                      + ' end'
                                 WHEN xtype IN ( 58, 61 )
                                 THEN 'case when [' + name
                                      + '] is null then ''NULL'' else '
                                      + ''''''''' + ' + 'CONVERT(varchar,['
                                      + name + '],121)' + '+''''''''' + ' end'
                                 WHEN xtype IN ( 40, 41, 42 )
                                 THEN 'case when [' + name
                                      + '] is null then ''NULL'' else '
                                      + ''''''''' + ' + 'cast([' + name
                                      + '] as varchar)' + '+''''''''' + ' end'
                                 WHEN xtype IN ( 167, 36 )
                                 THEN 'case when [' + name
                                      + '] is null then ''NULL'' else '
                                      + ''''''''' + ' + 'replace([' + name
                                      + '],'''''''','''''''''''')'
                                      + '+''''''''' + ' end'
                                 WHEN xtype IN ( 231 )
                                 THEN 'case when [' + name
                                      + '] is null then ''NULL'' else '
                                      + '''N'''''' + ' + 'replace([' + name
                                      + '],'''''''','''''''''''')'
                                      + '+''''''''' + ' end'
                                 WHEN xtype IN ( 175 )
                                 THEN 'case when [' + name
                                      + '] is null then ''NULL'' else '
                                      + ''''''''' + ' + 'cast(replace(['
                                      + name
                                      + '],'''''''','''''''''''') as Char('
                                      + CAST(length AS VARCHAR)
                                      + '))+''''''''' + ' end'
                                 WHEN xtype IN ( 239 )
                                 THEN 'case when [' + name
                                      + '] is null then ''NULL'' else '
                                      + '''N'''''' + ' + 'cast(replace(['
                                      + name
                                      + '],'''''''','''''''''''') as Char('
                                      + CAST(length AS VARCHAR)
                                      + '))+''''''''' + ' end'
                                 ELSE '''NULL'''
                            END AS Cols ,
                            name
                  FROM      syscolumns
                  WHERE     id = OBJECT_ID(@tablename)
                ) T
        IF ( @number != 0
             AND @number IS NOT NULL
           )
            BEGIN
                SET @sql = 'select top ' + CAST(@number AS VARCHAR(6000))
                    + ' ''INSERT INTO [' + @tablename + ']' + LEFT(@sql,
                                                              LEN(@sql) - 1)
                    + ') ' + LEFT(@sqlValues, LEN(@sqlValues) - 4)
                    + ')''  from ' + @tablename
            END
        ELSE
            BEGIN 
                SET @sql = 'select ''INSERT INTO [' + @tablename + ']'
                    + LEFT(@sql, LEN(@sql) - 1) + ') ' + LEFT(@sqlValues,
                                                              LEN(@sqlValues)
                                                              - 4)
                    + ')''  from ' + @tablename
            END

        IF ( @whereClause IS NOT NULL
             AND @whereClause <> ''
           )
            BEGIN
                SET @whereClause = ' where ' + REPLACE(@whereClause, '', '''')
            END 
            
           
        SET @sql = ' insert into #temp ([sql]) ' + @sql + @whereClause
		
        EXEC(@sql)
        
        SELECT  ' -- DELETE FROM ' + @tablename + ' ' + @whereClause [sql]
        UNION ALL
        SELECT  'IF NOT EXISTS ( SELECT 1 FROM ' + @tablename + ' '
                + @whereClause + ' ) '
        UNION ALL
        SELECT  'BEGIN'
        UNION ALL
        SELECT  *
        FROM    #temp
        UNION ALL
        SELECT  'END'
    END
GO`
	st, err := conn.Prepare(insertProcedure)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer st.Close()

	stmt1, err := conn.Prepare(`SELECT name FROM sys.tables WHERE name LIKE 'YGZ_%'`)

	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt1.Close()

	//通过Statement执行查询
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}

	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength) //临时存储每行数据
	for index, _ := range cache {              //为每一列初始化一个指针
		var a interface{}
		cache[index] = &a
	}
	var list []map[string]interface{} //返回的切片
	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{}) //取实际类型
		}
		list = append(list, item)
	}
	defer rows.Close()

	fp, err := os.Create("E:/MSSQLBACKUP/bak.sql") // 如果文件已存在，会将文件清空。
	if err != nil {
		log.Fatal("文件创建失败。")
	}
	defer fp.Close()

	for _, data := range list {
		createStr, err := conn.Prepare(fmt.Sprintf(`sp_helptable %s`, data["name"]))
		if err != nil {
			log.Fatal("Query failed:", err.Error())
		}
		//通过Statement执行查询
		rows, err := createStr.Query()
		if err != nil {
			log.Fatal("Query failed:", err.Error())
		}

		columns, _ := rows.Columns()
		columnLength := len(columns)
		cache := make([]interface{}, columnLength) //临时存储每行数据
		for index, _ := range cache {              //为每一列初始化一个指针
			var a interface{}
			cache[index] = &a
		}
		var createList []map[string]interface{} //返回的切片
		for rows.Next() {
			_ = rows.Scan(cache...)

			item := make(map[string]interface{})
			for i, data := range cache {
				item[columns[i]] = *data.(*interface{}) //取实际类型
			}
			createList = append(createList, item)
		}
		defer rows.Close()
		for _, temp := range createList {
			_, _ = io.WriteString(fp, temp["s"].(string))
		}

		stmt, err := conn.Prepare(fmt.Sprintf(`SELECT count(*) FROM %s`, data["name"]))

		if err != nil {
			log.Fatal("Prepare failed:", err.Error())
		}
		defer stmt1.Close()

		recordNum := "EXEC spGenInsertSQL @tableName = N'data_dict', @number = N'1000', @whereClause = N'1=1'"
		insertStr, err := conn.Prepare(fmt.Sprintf(`spGenInsertSQL %s`, data["name"]))
		if err != nil {
			log.Fatal("Query failed:", err.Error())
		}
		//通过Statement执行查询
		insertRows, err := insertStr.Query()
		if err != nil {
			log.Fatal("Query failed:", err.Error())
		}

		insertColumns, _ := insertRows.Columns()
		insertColumnLength := len(insertColumns)
		insertCache := make([]interface{}, insertColumnLength) //临时存储每行数据
		for index, _ := range cache {                          //为每一列初始化一个指针
			var a interface{}
			insertCache[index] = &a
		}
		var insertList []map[string]interface{} //返回的切片
		for insertRows.Next() {
			_ = rows.Scan(cache...)

			item := make(map[string]interface{})
			for i, data := range cache {
				item[columns[i]] = *data.(*interface{}) //取实际类型
			}
			insertList = append(insertList, item)
		}
		defer insertRows.Close()
		for _, temp := range insertList {
			_, _ = io.WriteString(fp, temp["s"].(string))
		}

		//cmdStr := fmt.Sprintf("exec master..xp_cmdshell 'bcp \"sp_helptable %s\" %s.dbo.[%s] queryout E:/MSSQLBACKUP/%s.bcp -c -S %s -U %s -P %s '", data["name"], database, data["name"], data["name"], server, user, password)
		//cmdStr := fmt.Sprintf("exec master..xp_cmdshell 'bcp \"sp_helptable %s\" %s.dbo.[%s] queryout E:/MSSQLBACKUP/%s.bcp -c -S %s -U %s -P %s '", data["name"], database, data["name"], data["name"], server, user, password)
		//fmt.Println(cmdStr)
		//cmd := exec.Command(cmdStr)
		//// 执行命令，并返回结果
		//output, err := cmd.Output()
		//if err != nil {
		//	panic(err)
		//}
		//// 因为结果是字节数组，需要转换成string
		//fmt.Println(string(output))
	}

}

func ZipFile() bool {
	err := Zip(outDir, ZIPOutputPath)
	if err != nil {
		return false
	}
	return true
}

func Zip(srcDir string, zipFile string) error {
	dir, err := ioutil.ReadDir(srcDir)
	if err != nil {
		log.Fatalf("ioutil.ReadDir err:%v", err)
	}
	if len(dir) == 0 {
		log.Fatal("压缩目录为空")
	}
	_ = os.RemoveAll(zipFile)
	file, err := os.Create(zipFile)
	if err != nil {
		log.Fatal("创建压缩文件失败")
	}
	defer file.Close()

	archive := zip.NewWriter(file)
	defer archive.Close()

	filepath.Walk(srcDir, func(path string, info fs.FileInfo, err error) error {
		if path == srcDir {
			return nil
		}
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, srcDir+`\`)
		if info.IsDir() {
			header.Name += `\`
		} else {
			header.Method = zip.Deflate
		}

		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})
	return nil
}
