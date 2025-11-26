export function formatTimestamp(timestamp: number) {
  if (!timestamp) {
    return ''
  }
  const date = new Date(timestamp)
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0') // Months are 0-based
  const day = String(date.getDate()).padStart(2, '0')
  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')
  const seconds = String(date.getSeconds()).padStart(2, '0')
  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
}

export function fromBirthToAge(timestamp: number): number | string {
  // Check for invalid or zero timestamp
  // if (!timestamp || timestamp <= 0) {
  //   return '未知'
  // }

  // Convert Unix timestamp (seconds) to Date object
  const birthDate = new Date(timestamp * 1000)

  // Get current date (October 5, 2025, based on provided context)
  const today = new Date() // JST timezone

  // Check if birthDate is valid
  if (Number.isNaN(birthDate.getTime())) {
    return '未知'
  }

  // Calculate age
  let age = today.getFullYear() - birthDate.getFullYear()
  const monthDiff = today.getMonth() - birthDate.getMonth()
  const dayDiff = today.getDate() - birthDate.getDate()

  // Adjust age if birthday hasn't occurred this year
  if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
    age--
  }

  // Ensure age is non-negative
  if (age < 0) {
    return '未知'
  }

  return age
}

export function formatArray(data: string | JSON) {
  if (typeof data === 'string') {
    if (data[0] === '[') {
      try {
        return JSON.parse(data)
      }
      catch (e: unknown) {
        console.error('格式化数组失败:', e, data)
        return []
      }
    }
    return data.split(',')
  }
  return data
}

export function getToken() {
  return localStorage.getItem('token') ?? ''
}

export function setToken(token: string) {
  return localStorage.setItem('token', token)
}
