// 设备API测试脚本
// 使用方法：在浏览器控制台中运行

// 测试创建设备
async function testCreateDevice() {
  try {
    const deviceData = {
      eqName: "测试设备",
      productsId: 1,
      eqInfo: "这是一个测试设备",
      status: "启用"
    };
    
    const response = await fetch('/api/wlEquipment/createWlEquipment', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + localStorage.getItem('token') // 需要登录token
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
    return result;
  } catch (error) {
    console.error('获取设备列表失败:', error);
  }
}

// 运行测试
console.log('开始测试设备API...');
// testCreateDevice();
// testGetDeviceList(); 