<template>
<div class="start" v-if="start == false && finish == false">
    <el-container class="starttitle">
        <div id="titletext">
        <span><font face="微软雅黑" size="10" color="#fff">准备好了吗？点击开始答题！</font></span>
        </div>
    </el-container>
    <el-container class="startbutton">
        <el-button
        id="startbutton"
        color="#cc9933"
        type="info"
        @click="start = true;initQues()">
            开始答题
        </el-button>
    </el-container>
</div>

<div v-loading="loading" class="test" v-if="start == true && finish == false">
    <el-container class="step">
        <el-progress :percentage="per" :stroke-width="20">
            {{question_form.progress}}
        </el-progress>
  </el-container>
  <el-container class="title">
      <span><br/><br/><font face="微软雅黑" size="6" color="#696969">{{question_form.desc}}</font></span>
  </el-container>
  <el-container class="options">
    <el-radio-group v-model="option" size="large">
        <button class="button"
            label="A" 
            @click="select('A')"
            >{{question_form.option_a}}</button>
    </el-radio-group>
  </el-container>
  <el-container class="options">
    <el-radio-group v-model="option" size="large">
        <button class="button"
            label="B" 
            @click="select('B')"
            >{{question_form.option_b}}</button>
    </el-radio-group>
  </el-container>
</div>

</template>

<script>
import { reactive, ref, getCurrentInstance } from 'vue'
import { useRouter } from 'vue-router'
import qs from "qs"
import { refAutoReset } from '@vueuse/shared'

export default ({
    setup() {
        const router = useRouter()
        const active = ref(0)
        const button_select = ref(false)
        const start = ref(false)
        const index = ref(0)
        const per = ref(0)
        let finish = ref(false)
        const loading = ref(true)

        

        const {proxy} = getCurrentInstance()

        let result_form = reactive({
            E: 0,
            I: 0,
            S: 0,
            N: 0,
            T: 0,
            F: 0,
            J: 0,
            P: 0,
        })

        let result = reactive({
            data: '',
        })

        let selection = []

        let getResult=()=> {
            if (result_form.E > result_form.I) {
                result.data = result.data  + 'E'
            } else {
                result.data = result.data + 'I'
            }

            if (result_form.S > result_form.N) {
                result.data  = result.data  + 'S'
            } else {
                result.data  = result.data  + 'N'
            }

            if (result_form.T > result_form.F) {
                result.data  = result.data  + 'T'
            } else {
                result.data  = result.data  + 'F'
            }

            if (result_form.J > result_form.P) {
                result.data  = result.data  + 'J'
            } else {
                result.data  = result.data  + 'P'
            }
        }

        let queslist = []

        let question_form = reactive({
            id: '',
            desc: '',
            option_a: '',
            option_b: '', 
            progress: 0,
        })

        const GetQues=async()=> {
            const listres = await new proxy.$request(proxy.$urls.m().getquestionidlist).get()
            const idList = listres.question_id_list
            //console.log(idList.length)
            // const obj = {"id":idList[index.value]}
            // const res = await new proxy.$request(proxy.$urls.m().getquestion, obj).get()

            for (let i=0; i<idList.length; i++) {
                const obj = {"id":idList[i]}
                const res = await new proxy.$request(proxy.$urls.m().getquestion, obj).get()
               // console.log(res)
                queslist.push({
                    id: res.question_info.question_id,
                    desc: res.question_info.question_desc,
                    option_a: res.question_info.option_a_desc,
                    option_b: res.question_info.option_b_desc,
                    option_a_target: res.question_info.option_a_target,
                    option_b_target: res.question_info.option_b_target,
                })
            }

            question_form.id = queslist[0].id
            question_form.desc = queslist[0].desc
            question_form.option_a = queslist[0].option_a
            question_form.option_b = queslist[0].option_b
            question_form.progress = '0' + '/' + queslist.length

           loading.value = false
        }

        GetQues()


        // //初始化第一道题目
        // const initQues=async()=> {
        //     const listres = await new proxy.$request(proxy.$urls.m().getquestionidlist).get()
        //     const idList = listres.question_id_list
        //     const obj = {"id":idList[index.value]}
        //     const res = await new proxy.$request(proxy.$urls.m().getquestion, obj).get()
        //     question_form.desc = res.question_info.question_desc
        //     question_form.option_a = res.question_info.option_a_desc
        //     question_form.option_b = res.question_info.option_b_desc
        //     question_form.progress = '0' + '/' + idList.length
        // }

        
        // 选择题目
        const select=async(option)=> {
            //根据索引index获取当前题目的信息
            // const listres = await new proxy.$request(proxy.$urls.m().getquestionidlist).get()
            // const idList = listres.question_id_list
            // let obj = {"id":idList[index.value]}
            // let res = await new proxy.$request(proxy.$urls.m().getquestion, obj).get()
            if (option == 'A') {
                selection.push({"id":queslist[index.value].id, "desc":queslist[index.value].desc, "option":option, "option_desc": queslist[index.value].option_a})
            } else if (option == 'B') {
                selection.push({"id":queslist[index.value].id, "desc":queslist[index.value].desc, "option":option, "option_desc": queslist[index.value].option_b})
            }

            // 分析传进的选项option A or B
            let target
            if (option == 'A') {
                target = queslist[index.value].option_a_target
            } else if (option == 'B') {
                target = queslist[index.value].option_b_target
            }
            switch (target) {
                case 'E':
                    result_form.E++
                    break;
                case 'I':
                    result_form.I++
                    break;
                case 'S':
                    result_form.S++
                    break;
                case 'N':
                    result_form.N++
                    break;
                case 'T':
                    result_form.T++
                    break;
                case 'F':
                    result_form.F++
                    break;
                case 'J':
                    result_form.J++
                    break;
                case 'P':
                    result_form.P++
                    break;
                default:
                    break;
            }

            //判断是否有下一题，如果有则初始化下一题
            const num = queslist.length //获取题目数量
            //如果没有下一道题了（一共n题，索引为n-1则尽头）
            if (index.value == num-1) {
                per.value = 100
                getResult()
                finish.value = true
                goResult(result.data) //去往结果页
                return
            }

            //如果还有下一道题，初始化下一道题目的信息
            index.value++
            // obj = {"id":idList[index.value]}
            // res = await new proxy.$request(proxy.$urls.m().getquestion, obj).get()
            question_form.desc = queslist[index.value].desc
            question_form.option_a = queslist[index.value].option_a
            question_form.option_b = queslist[index.value].option_b
            question_form.progress = index.value + '/' + num
            per.value = (index.value/num)*100
        }

        const goResult=async(param)=> {
            // console.log(toString(selection))
            const token = await new proxy.$request(proxy.$urls.m().verifytoken).get()
            const token_obj = {"username": token.username}
            const userres= await new proxy.$request(proxy.$urls.m().getuserinfo, token_obj).get()
            const obj = {"username":token.username, "type":param, "student_id":userres.data.userinfo.binding_student_id, "selection": JSON.stringify(selection)}
            await new proxy.$request(proxy.$urls.m().addtestdata, JSON.stringify(obj)).post()
            router.push('/result')
        }

        return {
            select,

            loading,
            GetQues,
           // initQues,
            per,
            index,
            start,
            finish,
            question_form,
            result_form,
            result,
            button_select, 
            active, 
            goResult,
            }
    },
})
</script>

<style scoped>
.start {
    background-color: #fff;
    height: 100%;
    padding-inline: 0; 
}
.starttitle {
    justify-content: center;
    height: 50%;
    background-color: #4298B4;
    width: 100%;
}
.starttitle #titletext {
    position: relative;
    top: 90px;
}
.startbutton {
    justify-content: center;
    height: 50%;
    background-color: #4298B4;
    width: 100%;
}

.startbutton .el-button {
    position: relative;
    top: 20px;
    color: #000;
    font-size: 20px;
    height: 70px;
    width: 25%;
    border-radius: 50px
}
.test {
    justify-content: center;
    background-color: #fff;
    height: 100%;
    padding-inline: 0;
}
.step {
        justify-content: center;
        background-color: #fff;
        position: relative;
        top: 20px;
        height: 10%;
        padding: 0%;
}

.el-progress--line {
    margin-top: 15px;
    margin-bottom: 15px;
    width: 35%;
}

.el-radio-group {
    justify-content: center;
    width: 100%;
}

#startbutton {
    color: #fff;
    font-size: 30px;
    height: 70px;
    width: 20%;
}

.button {
    background-color: #9999CC; /* Green */
    border: #fff;
    border-radius: 12px;
    color: white;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    height: 70px;
    width: 35%;
    font-size: 20px;
}

.title {
    justify-content: center;
    background-color: #fff;
    height: 30%;
    position: relative;
    top: 15px;
    padding: 0%; 
}

.options {
    justify-content: center;
    background-color: #fff;
    height: 20%;
    width: 100%;
    padding: 0%;
}

.end {
    background-color: #fff;
    height: 100%;
    padding-inline: 0; 
}

.endtitle {
    justify-content: center;
    height: 50%;
    background-color: #4298B4;
    width: 100%;
}
.endtitle #endtext {
    position: relative;
    top: 90px;
}
.endcontent {
    justify-content: center;
    height: 50%;
    background-color: #4298B4;
    width: 100%;
}
</style>
