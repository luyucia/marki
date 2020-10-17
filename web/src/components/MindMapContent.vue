<template>
    <div>
        <h1>思维导图</h1>
        <div>{{content}}</div>
        <p>aaa</p>
        <div id="chart"></div>
    </div>

</template>

<script>
    import * as d3 from "d3";

    export default {
        name: 'MindMapContent',
        props: {
            content: String
        },
        mounted() {
            let data = [4, 8, 15, 16, 23, 42]
            const div = d3.select("#chart")
                .style("font", "10px sans-serif")
                .style("text-align", "right")
                .style("color", "white");

            div.selectAll("div")
                .data(data)
                .join("div")
                .style("background", "steelblue")
                .style("padding", "3px")
                .style("margin", "1px")
                .style("width", d => `${d * 10}px`)
                .text(d => d);
        }
    }

    // const div = d3.create("div");
    // div.html("Hello, world!");


    //    https://github.com/d3/d3/wiki

    // let data = [4, 8, 15, 16, 23, 42]


</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

    .markdown_show {
        margin-left: auto;
        margin-right: auto;
        padding-top: 10px;
        width: 70%;
    }


</style>
