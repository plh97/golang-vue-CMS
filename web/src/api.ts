import { FMessage } from '@fesjs/fes-design'
import { getToken, setToken } from './common/utils'
import { baseURL } from './config'

export async function request(url: string, data: any = {}, options: any = {}) {
  const config: RequestInit = {
    method: options.method || 'GET',
    headers: {
      'Token': getToken() ?? '',
      'Content-Type': 'application/json',
      ...(options.headers ?? {}),
    },
  }
  // 处理请求体
  if (options.method && options.method !== 'GET') {
    if (data instanceof FormData) {
      config.body = data
      // eslint-disable-next-line ts/ban-ts-comment
      // @ts-ignore
      delete config.headers['Content-Type']
    }
    else {
      // eslint-disable-next-line ts/ban-ts-comment
      // @ts-ignore
      config.headers['Content-Type'] = 'application/json'
      config.body = JSON.stringify(data)
    }
  }
  return fetch(baseURL + url, config).then(async (response) => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    return response.json()
  }).then((data) => {
    // 处理响应内容异常
    if (data?.code === 2000) {
      FMessage.error({
        content: data?.msg,
      })
      setToken('')
      location.replace('/login')
      return Promise.reject(new Error('未登录或登录已过期，请重新登录'))
    }
    if (data?.code !== 0) {
      // Reject the promise with an error object containing code and message
      FMessage.error({
        content: data?.msg,
      })
      return Promise.reject(new Error(data?.msg))
    }
    if (options.transformData) {
      return options.transformData(data)
    }
    return data.data
  })
}
