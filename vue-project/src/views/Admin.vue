<template>
    <div class="header">
        <div class="searchbox">
            <div class="search">
                <a-input-search
                v-model:value="searchUsername"
                placeholder="input username to search"
                enter-button="Search"
                size="large"
                @search="onSearch"
                />
                <a-button type="primary" size="large" @click="getUserList">Select All</a-button>
            </div> 
            <a  href="/" style="position: absolute; top: 25px; left: 90%; margin-left:50px; font-size: large;">Sign out</a>
        </div>
    </div>
    <div class="Listbox">  
        <a-table :columns="columns" :row-key="id" :data-source="userlist" :pagination="pagination" >
        </a-table>
    </div>
    <div class="operationbox">
        <div class="operation1">
        <a-form
        :model="formdata"
        >
        <a-input
        v-model:value="updateId"
        placeholder="input id"
        style="width: 100px"
        />
        <a-input
        v-model:value="formdata.username"
        placeholder="input username"
        style="width: 200px"
        />
        <a-input
        v-model:value="formdata.password"
        placeholder="input password"
        style="width: 200px"
        />
        <a-button type="primary" html-type="submit" @click="update">Update</a-button>
        </a-form>
    </div>
    <div class="operation2">
        <a-input
        v-model:value="deleteId"
        placeholder="input id"
        style="width: 200px"
        />
        <a-button type="primary" html-type="submit" @click="deleteUser">Delete</a-button>
    </div>
    
    </div>
</template>

<style scoped>
    .header {
        display: flex;
        width: 100%;
        background-color: rgb(189, 234, 248);
        height: 80px;
    }

    .Listbox {
        position: absolute;
        top: 15%;
        left: 20%;
        width: 1000px;
        height: 100%;
    }

    .search {
        position: relative;
        display: flex;
        top: 25%;
        left: 80%;
        width: 600px;
    }
    .operationbox {
        position: absolute;
        top: 66%;
        left: 19%;
    }

    .operation1 {
        margin: 20px;
    }

    .operation2 {
        position: absolute;
        top: 75%;
        margin-left: 20px;
        margin-top: 10px;
    }
</style>

<script>

import admin from './Login.vue'

console.log(admin.username)

const columns = [
    {
    title: 'ID',
    dataIndex: 'id',
    width: '10%',
    key:'id'
  },
  {
    title: 'Username',
    dataIndex: 'username',
    width: '20%',
    key:'username'
  },
  {
    title: 'Password',
    dataIndex: 'password',
    width: '20%',
    key:'password'
  },
]

export default {
    data() {
        return {
            searchUsername: '',
            updateId: '',
            formdata: {
                username: '',
                password: ''
            },
            userlist: [],
            columns,
            deleteId:''
        }
    },
    created() {
        this.getUserList()
    },
    methods: {
        async getUserList() {
            const {data: res} = await this.$http.get('/users')
            console.log(res.data)
            this.userlist = res.data
        },
        async update() {
            console.log(this.formdata)
            const {data: res} = await this.$http.post('/user/'+this.updateId, this.formdata)
            console.log(res)
            this.getUserList()
        },
        async deleteUser() {
            const {data: res} = await this.$http.delete('/user/'+this.deleteId)
            console.log(this.deleteId)
            this.getUserList()
        },
        async onSearch() {
            const {data: res} = await this.$http.get('/user/'+this.searchUsername)
            console.log(res)
            this.userlist = res.data
        },
    }
}
</script>