package com.xidian.grades;

import java.io.IOException;
import org.apache.hadoop.fs.Path;
import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.lib.input.FileInputFormat;
import org.apache.hadoop.mapreduce.lib.output.FileOutputFormat;
import org.apache.hadoop.io.DoubleWritable;
import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.mapreduce.Job;
import org.apache.hadoop.mapreduce.Mapper;
import org.apache.hadoop.mapreduce.Reducer;
import org.apache.hadoop.util.GenericOptionsParser;

public class ClassAverageGrades {
    static class MyMapper extends Mapper<Object, Text, Text, IntWritable> {
        @Override
        protected void map(Object key, Text value, Mapper<Object, Text, Text, IntWritable>.Context context)
                throws IOException, InterruptedException {
            String[] strs = value.toString().split(",");
            context.write(new Text(strs[0] + "-" + strs[2]), new IntWritable(Integer.valueOf(strs[4])));
        }
    }

    static class MyReducer extends Reducer<Text, IntWritable, Text, DoubleWritable> {
        @Override
        protected void reduce(Text key, Iterable<IntWritable> values,
                Reducer<Text, IntWritable, Text, DoubleWritable>.Context context)
                throws IOException, InterruptedException {
            int sum = 0, num = 0;
            for (IntWritable i : values) {
                sum += i.get();
                num++;
            }
            context.write(key, new DoubleWritable((double) sum / num));
        }
    }

    public static void main(String[] args) throws IOException, ClassNotFoundException, InterruptedException {
        Configuration conf = new Configuration();
        String[] otherArgs = new GenericOptionsParser(conf, args).getRemainingArgs();
        if (otherArgs.length < 2) {
            System.err.println("Usage ClassAverageGrades <in> [<in>...] <out>");
            System.exit(-1);
        }

        Job job = Job.getInstance(conf, "class-average-grades");

        job.setJarByClass(StudentAverageGrades.class);

        job.setMapperClass(MyMapper.class);
        job.setReducerClass(MyReducer.class);
        
        job.setMapOutputKeyClass(Text.class);
        job.setMapOutputValueClass(IntWritable.class);
        
        job.setOutputKeyClass(Text.class);
        job.setOutputValueClass(DoubleWritable.class);
        
        
        for (int i = 0; i < otherArgs.length - 1; i++) {
            FileInputFormat.addInputPath(job, new Path(otherArgs[i]));
        }

        FileOutputFormat.setOutputPath(job, new Path(otherArgs[args.length - 1]));

        System.exit(job.waitForCompletion(true) ? 0 : 1);
    }
}
