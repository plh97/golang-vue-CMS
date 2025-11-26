export function validEmail(email: string) {
  // Check if email is a string and not empty
  if (typeof email !== 'string' || !email.trim()) {
    return false
  }
  // Regular expression for email validation
  // eslint-disable-next-line regexp/no-unused-capturing-group
  const emailRegex = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\])|(([a-z\-0-9]+\.)+[a-z]{2,}))$/i
  return emailRegex.test(email.trim())
}

export function useValidator(data: Record<string, any>) {
  function validator(v: any) {
    let field = v.field
    if (v.type === 'array') {
      const fieldList = JSON.parse(field)
      for (field of fieldList) {
        const val = data[field]
        if (!val) {
          return false
        }
      }
      return true
    }
    const val = data[field]
    if (v.type === 'email') {
      return validEmail(val)
    }
    if (typeof val === 'number') {
      return !!val
    }
    return !!val?.length
  }
  return validator
}
