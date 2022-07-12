<template>
<div class="container">
    
    <div class="loginbox">
        <div class="logintag">
            Sign up
        </div>
        <a-form
        :model="formdata"
        style="margin-left: 70px;margin-top: 80px;"
        name="basic"
        :label-col="{ span: 8 }"
        :wrapper-col="{ span: 16 }"
        autocomplete="off"
        >
            <a-form-item
            style="margin-bottom: 50px;"
            label="Username"
            name="username"
            >
            <a-input v-model:value="formdata.username"/>
            </a-form-item>

            <a-form-item
            style="margin-bottom: 50px;"
            label="Password"
            >
            <a-input-password  v-model:value="formdata.password"/>
            </a-form-item>

            <a-form-item :wrapper-col="{ offset: 5, span: 16 }">
            <a-button type="primary" html-type="submit" @click="register">Submit</a-button>
            </a-form-item>
            <a href="/Login">Return to sign in</a>
        </a-form>
    </div>
</div>  
</template>

<script>
export default {
    data() {
        return {
            formdata: {
                username: '',
                password: ''
            }
        }
    },
    methods: {
        async register() {
                const res = await this.$http.post('/register',this.formdata)
                console.log(this.formdata)
                console.log(res.data.code) 
                alert("SUCSESSFUL！");
                this.$router.push('Login')
        }
    }
}
</script>

<style scoped>
    .container {
        height: 100%;
        background-color: rgb(206, 206, 206);
    }

    .loginbox {
        display: flex;
        justify-content: flex;
        width: 450px;
        height: 450px;
        background-color: #fff;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        border-radius: 9px;
        align-items: center;
        
    }

    .logintag {
        position: absolute;
        top: 40px;
        left: 35%;
        font-size: 40px;
        border-radius: 9px;
    }
</style>
