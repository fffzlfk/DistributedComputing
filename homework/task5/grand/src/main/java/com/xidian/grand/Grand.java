package com.xidian.grand;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

import org.apache.hadoop.fs.Path;
import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Job;
import org.apache.hadoop.mapreduce.Mapper;
import org.apache.hadoop.mapreduce.Reducer;
import org.apache.hadoop.util.GenericOptionsParser;
import org.apache.hadoop.mapreduce.lib.input.FileInputFormat;
import org.apache.hadoop.mapreduce.lib.output.FileOutputFormat;


public class Grand {
    static class MyMapper extends Mapper<Object, Text, Text, Text> {
        @Override
        protected void map(Object key, Text value, Mapper<Object, Text, Text, Text>.Context context)
                throws IOException, InterruptedException {
            String[] strs = value.toString().split(",");
            String child = strs[0].trim();
            String parent = strs[1].trim();
            context.write(new Text(child), new Text("child-" + child + "-" + parent));
            context.write(new Text(parent), new Text("parent-" + child + "-" + parent));
        }
    }

    static class MyReducer extends Reducer<Text, Text, Text, Text> {
        @Override
        protected void reduce(Text key, Iterable<Text> values, Reducer<Text, Text, Text, Text>.Context context)
                throws IOException, InterruptedException {
            List<String> childs = new ArrayList<>();
            List<String> parents = new ArrayList<>();
            for (Text text : values) {
                String[] strs = text.toString().split("-");
                if (strs[0].equals("child")) {
                    childs.add(strs[2]);
                } else if (strs[0].equals("parent")) {
                    parents.add(strs[1]);
                }
            }
            for (String grandchild : parents) {
                for (String grandparent : childs) {
                    context.write(new Text(grandchild), new Text(grandparent));
                }
            }
        }
    }

    public static void main(String[] args) throws IOException, ClassNotFoundException, InterruptedException {
        Configuration conf = new Configuration();
        String[] otherArgs = new GenericOptionsParser(conf, args).getRemainingArgs();
        if (otherArgs.length < 2) {
            System.err.println("Usage Grand <in> [<in>...] <out>");
            System.exit(-1);
        }

        Job job = Job.getInstance(conf, "grand");

        job.setJarByClass(Grand.class);

        job.setMapperClass(MyMapper.class);
        job.setReducerClass(MyReducer.class);

        job.setMapOutputKeyClass(Text.class);
        job.setMapOutputValueClass(Text.class);

        job.setOutputKeyClass(Text.class);
        job.setOutputValueClass(Text.class);

        for (int i = 0; i < otherArgs.length - 1; i++) {
            FileInputFormat.addInputPath(job, new Path(otherArgs[i]));
        }

        FileOutputFormat.setOutputPath(job, new Path(otherArgs[args.length - 1]));

        System.exit(job.waitForCompletion(true) ? 0 : 1);
    }
}
