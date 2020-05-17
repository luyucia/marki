<template>
    <div class="marki_menu">
      <template v-for="(sub,i) in data" >
        <el-menu-item v-if="sub.child==null" :key="sub.url" :index="fatherIndex+'-'+i+':'+sub.url">
          <span slot="title">{{sub.name}}</span>
        </el-menu-item>
        <el-submenu v-if="sub.child!=null" :key="sub.name" :index="fatherIndex+'-'+i">
          <template slot="title">
            <span>{{sub.name}}</span>
          </template>
          <Menu :data="sub.child" :father-index="fatherIndex+'-'+i"></Menu>
        </el-submenu>
      </template>
    </div>
</template>

<script>
    export default {
        name: 'Menu',
        props: ['data','fatherIndex'],
      // ,{
      //       data: [],
      //       fatherIndex: Number
      //   },
        data: function () {
            return {
                menu: {}
            }
        },
        mounted() {
            // let app = this
            // this.$http.get('/get_menu').then(response =>{
            //   console.log(response.body)
            //   app.menu = response.body.data
            // })

        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    /*h3 {*/
    /*  margin: 40px 0 0;*/
    /*}*/
    /*ul {*/
    /*  list-style-type: none;*/
    /*  padding: 0;*/
    /*}*/
    /*li {*/
    /*  display: inline-block;*/
    /*  margin: 0 10px;*/
    /*}*/
    /*a {*/
    /*  color: #42b983;*/
    /*}*/
</style>
