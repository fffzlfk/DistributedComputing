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


if __name__ == '__main__':
    conf = SparkConf().setMaster("local").setAppName("averagegrades")
    sc = SparkContext(conf=conf)
    data = sc.textFile("/input/grades.txt").filter(filter)
    res = data.map(mapper).reduceByKey(reducer).map(
        lambda x: (x[0], '%.3f' % (x[1][1]/x[1][0])))
    res.saveAsTextFile("/output/task3/1")
