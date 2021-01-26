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
##### Create
```
/* Set primary keys and foreign keys */
CREATE TABLE Student (
	id INT,
	name VARCHAR(30),
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

##### Insert
```sql
INSERT INTO Student 
VALUES (1, "Kevin", 2021), (2, "Jenny", 2022)
```