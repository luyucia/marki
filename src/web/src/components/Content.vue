<template>
    <div v-html="content" class="markdown_show"></div>
</template>

<script>
    export default {
        name: 'MarkContent',
        props: {
            content: String
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

    .markdown_show {
        margin-left: auto;
        margin-right: auto;
        padding-top: 10px;
        width: 70%;
    }

    /*#app .el-row .markdown_show h1, h2, h3, h4, h5, h6 {*/
    /*    font-family: "TimesNewRomanPS-ItalicMT", "PingFang SC", "Lantinghei SC", "Microsoft Yahei", "Hiragino Sans GB", "Microsoft Sans Serif", "WenQuanYi Micro Hei", sans, serif;*/
    /*    color: #42b983;*/
    /*}*/

    /*#app .markdown_show code {*/
    /*    background-color: #42b983;*/
    /*}*/

    /*h3 {*/
    /*    color: #42b983;*/
    /*}*/


</style>
