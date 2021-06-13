mvn package
hadoop jar target/grades-1.0-SNAPSHOT.jar com.xidian.grades.StudentAverageGrades /input/grades.txt /output/task1/1
hadoop jar target/grades-1.0-SNAPSHOT.jar com.xidian.grades.ClassAverageGrades /input/grades.txt /output/task1/2