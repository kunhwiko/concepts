# SQL 
### Basics
---
##### Terminology 
```
Schema: name of the table and column (attribute) names 
Tuple: rows/records 
Degree/Arity: number of elements
Cardinality: number of rows 

Key : subset of columns that uniquely defines each record in a table 
Foreign Key : a reference to a key in another table 
```

##### Initial Setting
```
Not Null : values in the column cannot be null 
Primary Key : values in the column must be unique, specifies the key of the table, does not allow null values  
Autoincrement : unique number if automatically generated through incrementation for new records 
Unique : ensures values in a column are unique, many columns can be unique, allow null values  
```

<br />

### Creating Databases
---
##### Setting Keys in Create Clause
```sql
/* Set primary keys and foreign keys */
CREATE TABLE Student (
    id INT,
    name VARCHAR(30),
    school VARCHAR(60),
    graduation DATE,
    PRIMARY KEY (id)
)

/* Primary keys can be multiple columns if they make up the unique identifer, 
   Here, building_no and room_no together make up one unique identifier for each tuple */
CREATE TABLE Housing (
    building_no INT,
    room_no INT,
    resident_id INT,
    building_name VARCHAR(30),
    PRIMARY KEY (building_no, room_no),
    FOREIGN KEY resident_id REFERENCES Student (id)
)
```

<br />

### Keywords 
---
##### DISTINCT
```sql
/* Returns unique values */
SELECT DISTINCT name
FROM Student
WHERE school = 'University of Pennsylvania'
```

##### INTO
```sql
/* create a temporary table temp for the duration of the program */
SELECT id INTO temp
FROM Student
WHERE school = 'University of Pennsylvania'
```

##### WITH
```sql
WITH PennStudent AS (
    SELECT *
    FROM Student
    WHERE school = 'University of Pennsylvania'
)
SELECT P.name, P.graduation, H.building_name 
FROM PennStudent P 
JOIN Housing H ON P.id = H.resident_id 
```

<br />

### Set Operations 
---
##### UNION / ALL 
```sql
/* 
    schema must be the same for set operations 
    since sets do not return duplicate values
    we specify all to retain duplicates
*/

(SELECT name FROM Student)
UNION ALL 
(SELECT name FROM Professor)
```

##### Intersect 
```sql
(SELECT name FROM Student)
INTERSECT 
(SELECT name FROM Professor)
```

<br />

### Boolean Operations 
---
##### IN / NOT IN 
```sql
SELECT *
FROM Customers 
WHERE Country IN (
    SELECT Country 
    FROM Suppliers
)
```

##### ALL / ANY 
```sql
/* Return student_ids whose expected grades are
   higher or equal to everyone in Takes  */
SELECT DISTINCT student_id 
FROM Takes 
WHERE expected_grade >= ALL (
    SELECT expected_grade 
    FROM Takes 
)

/* Return student_ids whose expected grades are 
   at least higher than someone in Takes  */
SELECT DISTINCT student_id 
FROM Takes 
WHERE expected_grade > ANY (
    SELECT expected_grade 
    FROM Takes 
)
```

##### EXISTS / NOT EXISTS 
```sql
/* check the existance of any record in a subquery */
SELECT supplier_name 
FROM Suppliers s
WHERE EXISTS (
    SELECT product_name 
    FROM Products p  
    WHERE p.supplier_id = s.supplier_id 
)
```

<br />

### Aggregate Operations 
---
##### HAVING 
```sql
/* HAVING is to GROUP BY, WHERE is to FROM */
SELECT subject, COUNT(*) as num 
FROM Course
WHERE subject != 'AI'
GROUP BY subject 
HAVING COUNT(*) > 1
```

##### NULL 
```
Three State Logic 
| AND           | OR            | NOT         |
| ------------- |:-------------:| ----------: |
| T AND T --> T | T OR U --> T  | NOT U --> U |
| F AND U --> F | F OR F --> F  |             |
| T AND U --> U | F OR U --> U  |             |
| U AND U --> U | U OR U --> U  |             |

1) if x = NULL, then x = "June" is UNKNOWN
2) if x = NULL, then 3-x is NULL
3) max(age) discards NULLs (if all values are NULL, then NULL)
4) count(id) includes NULLs
5) JOIN disregards matching NULLs
6) SELECT DISTINCT treats NULLs as equals (there will only be one NULL)
7) COUNT DISTINCT eliminates all NULLs before counting
8) GROUP BY treats NULLs as equals 
```

##### JOIN Types 
```
1) Inner Join: join on a particular key, both instances are kept 
2) Natural Join: join on a particular key, only one instance of the key is kept 
3) Left Outer Join: include tuples of left table even if there are no matches 
4) Right Outer Join: include tuples of right table even if there are no matches 
5) Full Outer Join: include tuples of both tables even if there are no matches    
```