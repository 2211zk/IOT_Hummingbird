// 模拟标准品类数据
export const mockCategories = [
  {
    id: 1,
    name: '电子产品',
    code: 'ELEC001',
    category: '电子设备',
    description: '各类电子产品和设备',
    status: 1,
    createdAt: '2024-01-01T00:00:00Z',
    updatedAt: '2024-01-01T00:00:00Z'
  },
  {
    id: 2,
    name: '家用电器',
    code: 'HOME001',
    category: '家居用品',
    description: '家庭使用的各类电器产品',
    status: 1,
    createdAt: '2024-01-02T00:00:00Z',
    updatedAt: '2024-01-02T00:00:00Z'
  },
  {
    id: 3,
    name: '办公用品',
    code: 'OFFICE001',
    category: '办公设备',
    description: '办公室使用的各类用品',
    status: 0,
    createdAt: '2024-01-03T00:00:00Z',
    updatedAt: '2024-01-03T00:00:00Z'
  },
  {
    id: 4,
    name: '服装鞋帽',
    code: 'CLOTH001',
    category: '服装配饰',
    description: '各类服装和配饰产品',
    status: 1,
    createdAt: '2024-01-04T00:00:00Z',
    updatedAt: '2024-01-04T00:00:00Z'
  },
  {
    id: 5,
    name: '食品饮料',
    code: 'FOOD001',
    category: '食品',
    description: '各类食品和饮料产品',
    status: 1,
    createdAt: '2024-01-05T00:00:00Z',
    updatedAt: '2024-01-05T00:00:00Z'
  }
]

// 模拟类别数据
export const mockCategoryTypes = [
  '电子设备',
  '家居用品',
  '办公设备',
  '服装配饰',
  '食品'
]

// 模拟API响应
export const mockApiResponse = {
  success: {
    code: 0,
    data: {
      list: mockCategories,
      total: mockCategories.length,
      page: 1,
      pageSize: 10
    },
    msg: '获取成功'
  },
  error: {
    code: 1,
    data: null,
    msg: '获取失败'
  },
  networkError: new Error('Network Error'),
  timeoutError: { code: 'ECONNABORTED', message: 'timeout' }
}

// 模拟分页数据
export const createMockPageData = (page = 1, pageSize = 10, total = 50) => {
  const start = (page - 1) * pageSize
  const end = Math.min(start + pageSize, total)
  
  const list = []
  for (let i = start; i < end; i++) {
    list.push({
      id: i + 1,
      name: `品类${i + 1}`,
      code: `CODE${String(i + 1).padStart(3, '0')}`,
      category: mockCategoryTypes[i % mockCategoryTypes.length],
      description: `这是品类${i + 1}的描述信息`,
      status: Math.random() > 0.2 ? 1 : 0,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString()
    })
  }
  
  return {
    code: 0,
    data: {
      list,
      total,
      page,
      pageSize
    },
    msg: '获取成功'
  }
}

// 模拟搜索结果
export const createMockSearchData = (keyword = '', category = '', status = null) => {
  let filteredData = [...mockCategories]
  
  if (keyword) {
    filteredData = filteredData.filter(item => 
      item.name.includes(keyword) || 
      item.code.includes(keyword) ||
      item.description.includes(keyword)
    )
  }
  
  if (category) {
    filteredData = filteredData.filter(item => item.category === category)
  }
  
  if (status !== null) {
    filteredData = filteredData.filter(item => item.status === status)
  }
  
  return {
    code: 0,
    data: {
      list: filteredData,
      total: filteredData.length,
      page: 1,
      pageSize: 10
    },
    msg: '获取成功'
  }
}