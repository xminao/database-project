//所有请求接口

const url = "http://localhost:8888"

const urls = class{
    static m(){
        const register = `${url}/user/register`
        const login = `${url}/user/login`
        const getuserinfo = `${url}/user/getuserinfo`

        return {
            register,
            login,
            getuserinfo
        }
    }

}
export default urls

