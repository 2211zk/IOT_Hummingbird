// 设备选择器测试脚本
// 使用方法：在浏览器控制台中运行

// 测试获取设备列表
async function testGetDeviceList() {
  try {
    const response = await fetch('/api/wlEquipment/getWlEquipmentList?page=1&pageSize=10', {
      method: 'GET',
      headers: {
        'Authorization': 'Bearer ' + localStorage.getItem('token')
      }
    });
    
    const result = await response.json();
    console.log('获取设备列表结果:', result);
    
    if (result.code === 0) {
      console.log('设备列表数据:', result.data.list);
      console.log('设备总数:', result.data.total);
    } else {
      console.error('获取设备列表失败:', result.msg);
    }
    
    return result;
  } catch (error) {
    console.error('获取设备列表失败:', error);
  }
}

// 测试创建设备（如果列表为空）
async function testCreateDevice() {
  try {
    const deviceData = {
      eqName: "测试设备001",
      productsId: 1,
      eqInfo: "这是一个测试设备",
      status: "启用"
    };
    
    const response = await fetch('/api/wlEquipment/createWlEquipment', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + localStorage.getItem('token')
      },
      body: JSON.stringify(deviceData)
    });
    
    const result = await response.json();
    console.log('创建设备结果:', result);
    return result;
  } catch (error) {
    console.error('创建设备失败:', error);
  }
}

// 测试部门设备关联
async function testDepartmentDeviceAssociation() {
  try {
    const departmentData = {
      name: "测试部门",
      leader: "测试负责人",
      phone: "13800138000",
      email: "test@example.com",
      status: "启用",
      deviceIds: [1] // 关联设备ID为1的设备
    };
    
    const response = await fetch('/api/department/create', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + localStorage.getItem('token')
      },
      body: JSON.stringify(departmentData)
    });
    
    const result = await response.json();
    console.log('创建部门结果:', result);
    return result;
  } catch (error) {
    console.error('创建部门失败:', error);
  }
}

// 运行测试
console.log('开始测试设备选择器功能...');

// 1. 首先检查是否有设备数据
testGetDeviceList().then(result => {
  if (result.code === 0 && result.data.total === 0) {
    console.log('设备列表为空，创建测试设备...');
    testCreateDevice().then(() => {
      console.log('测试设备创建完成，重新获取设备列表...');
      testGetDeviceList();
    });
  }
});

// 2. 测试部门设备关联
// testDepartmentDeviceAssociation(); 