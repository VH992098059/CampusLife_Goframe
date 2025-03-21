import { useState } from 'react';
import { Layout, Input, Button, List, Card, Typography } from 'antd';
import { MessageOutlined, SendOutlined } from '@ant-design/icons';

const { Content, Sider } = Layout;
const { TextArea } = Input;
const { Title, Paragraph } = Typography;

export const AiChat = () => {
  const [messages, setMessages] = useState([
    { sender: '客服', content: '您好！请问有什么可以帮到您的？', time: '15:30' }
  ]);
  const [inputMessage, setInputMessage] = useState('');

  const handleSend = () => {
    if (inputMessage.trim()) {
      setMessages([
        ...messages,
        {
          sender: '用户',
          content: inputMessage,
          time: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
        }
      ]);
      setInputMessage('');
    }
  };

  return (
    <Layout style={{ minHeight: '100vh' }}>
      {/* 左侧聊天侧边栏 */}
      <Sider
        width={300}
        theme="dark"
        style={{
          background: '#001529',
          padding: '20px',
          borderRight: '1px solid #1890ff'
        }}
      >
        <div style={{ color: 'white', marginBottom: 24 }}>
          <Title level={4} style={{ color: 'white' }}>
            <MessageOutlined /> 联系我们的客服
          </Title>
          <Paragraph type="secondary">上次活动 一小时前</Paragraph>
        </div>

        <div style={{ height: 'calc(100vh - 180px)', overflowY: 'auto' }}>
          <List
            dataSource={messages}
            renderItem={item => (
              <div
                style={{
                  background: item.sender === '客服' ? '#1890ff' : '#f0f0f0',
                  color: item.sender === '客服' ? 'white' : 'black',
                  padding: '8px 12px',
                  borderRadius: '8px',
                  marginBottom: '8px',
                  maxWidth: '80%',
                  float: item.sender === '客服' ? 'left' : 'right'
                }}
              >
                <div>{item.content}</div>
                <div style={{ fontSize: '0.75em', opacity: 0.8 }}>{item.time}</div>
              </div>
            )}
          />
        </div>

        <div style={{ marginTop: 'auto' }}>
          <TextArea
            value={inputMessage}
            onChange={(e) => setInputMessage(e.target.value)}
            placeholder="输入你的信息..."
            autoSize={{ minRows: 1, maxRows: 4 }}
            onPressEnter={(e) => {
              if (!e.shiftKey) {
                e.preventDefault();
                handleSend();
              }
            }}
          />
          <Button
            type="primary"
            icon={<SendOutlined />}
            onClick={handleSend}
            style={{ marginTop: 8, float: 'right' }}
          >
            发送
          </Button>
        </div>
      </Sider>

      {/* 右侧主内容区域 */}
      <Content style={{ padding: '24px', background: '#f0f2f5' }}>
        <Card title="欢迎使用SAILING服务" style={{ marginBottom: 24 }}>
          <Title level={4}>最新公告</Title>
          <Paragraph>
            <ul>
              <li>客户端v2.3.0已发布</li>
              <li>教程下载已更新</li>
              <li>官网地址：https://www.sailing.example</li>
            </ul>
          </Paragraph>
        </Card>

        <Card title="使用指南">
          <Title level={5}>快速开始</Title>
          <Paragraph>
            1. 访问官网下载客户端<br/>
            2. 在配置文件中设置您的API密钥<br/>
            3. 参考教程进行手动配置
          </Paragraph>
        </Card>
      </Content>
    </Layout>
  );
};
