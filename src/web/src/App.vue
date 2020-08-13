<template>
    <div id="app">
        <el-container style="height: 98vh;">
            <el-aside>
                <el-menu class="el-menu-vertical-demo" @select="loadContent">
                    <Menu :data="menu.child" :fatherIndex="0"></Menu>
                </el-menu>
            </el-aside>
            <el-main>
                <MarkContent v-if="content_type==='md' "  :content="content"></MarkContent>
                <MindMapContent v-if="content_type==='map' "  :content="content"></MindMapContent>
            </el-main>
        </el-container>
        <!--        <el-row :gutter="10">-->
        <!--            <el-col :xs="8" :sm="6" :md="4" :lg="4" :xl="5">-->

        <!--                <el-menu default-active="2" class="el-menu-vertical-demo" @select="loadContent">-->
        <!--                    <Menu :data="menu.child" :fatherIndex="0"></Menu>-->
        <!--                </el-menu>-->
        <!--            </el-col>-->
        <!--            <el-col :xs="4" :sm="6" :md="8" :lg="20" :xl="19">-->
        <!--                <MarkContent :content="content"></MarkContent>-->
        <!--            </el-col>-->
        <!--        </el-row>-->
    </div>
</template>

<script>
    import Menu from "./components/Menu";
    import MarkContent from "./components/MarkdownContent";
    import MindMapContent from "./components/MindMapContent";

    export default {
        name: 'App',
        components: {
            Menu,
            MarkContent,
            MindMapContent
        },
        data: function () {
            return {
                menu: {},
                content: "",
                content_type: "md"
            }
        },
        mounted() {
            let app = this
            this.$http.get('/get_menu').then(response => {
                app.menu = response.body.data
            })
        }
        , methods: {
            loadContent: function (index) {
                let url = index.split(':')[1]
                let app = this
                this.$http.get('/get_content?id=' + url).then(response => {
                    console.log(response.body.content_type)
                    console.log(response.body.content)
                    app.content = response.body.content
                    app.content_type = response.body.content_type
                })
            }
        }
    }
</script>

<style>
    #app {
        font-family: Avenir, Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: left;

    }


</style>
