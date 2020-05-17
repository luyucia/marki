<template>
    <div id="app">

        <el-row :gutter="10">
            <el-col :xs="8" :sm="6" :md="4" :lg="4" :xl="5">

                <el-menu default-active="2" class="el-menu-vertical-demo" @select="loadContent">
                    <Menu :data="menu.child" :fatherIndex="0"></Menu>
                </el-menu>
            </el-col>
            <el-col :xs="4" :sm="6" :md="8" :lg="20" :xl="19">
                <MarkContent :content="content"></MarkContent>
            </el-col>
        </el-row>
    </div>
</template>

<script>
    import Menu from "./components/Menu";
    import MarkContent from "./components/Content";

    export default {
        name: 'App',
        components: {
            Menu,
            MarkContent
        },
        data: function () {
            return {
                menu: {},
                content: ""
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
                    app.content = response.body
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
