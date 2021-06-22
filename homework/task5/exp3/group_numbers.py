# -*- coding: UTF-8 -*-

from pyspark import SparkConf, SparkContext


def filter(line):
    strs = line.split(',')
    return strs[3] == '必修'


def mapper(line):
    strs = line.split(',')
    return (strs[1], (1, int(strs[4])))


def reducer(x, y):
    return (x[0]+y[0], x[1]+y[1])


def classify(pair):
    key = ''
    if pair[1] >= 90:
        key = '90~100'
    elif pair[1] >= 80:
        key = '80~89'
    elif pair[1] >= 70:
        key = '70~79'
    elif pair[1] >= 60:
        key = '60~69'
    else:
        key = '60以下'
    return (key, pair[0])


if __name__ == '__main__':
    conf = SparkConf().setMaster("local").setAppName("numbers")
    sc = SparkContext(conf=conf)
    data = sc.textFile("/input/grades.txt").filter(filter)
    average_grades = data.map(mapper).reduceByKey(reducer).map(lambda x: (
        x[0], round(x[1][1]/x[1][0])))
    average_grades.saveAsTextFile("/output/task3/1")
    group_items = average_grades.map(classify).countByKey().items()
    sc.parallelize(group_items).saveAsTextFile("/output/task3/2")
