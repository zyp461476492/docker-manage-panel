import React from 'react';
import { connect } from 'dva';
import { Row, Col, Alert, Modal, Form, Input, Select, message } from 'antd';
import styles from './dialog.css';
import ContainerNetworkInput from '../containerNetworkInput/containerNetworkInput';

const InputGroup = Input.Group;
const { Option } = Select;
class ContainerFormPanel extends React.Component {
  render() {
    const { getFieldDecorator } = this.props.form;
    return (
      <Form layout="vertical">
        <Form.Item label="容器名称">
          {getFieldDecorator('containerName', {
            rules: [{ required: true, message: '请输入容器名称' }],
          })(<Input />)}
        </Form.Item>
        <Form.Item label="镜像名称">
          {getFieldDecorator('imageName', {
            rules: [{ required: true, message: '请输入镜像名称' }],
          })(<Input />)}
        </Form.Item>
        <div>
          <Form.Item label="端口绑定">
            {getFieldDecorator('ip', {
              initialValue: { dockerPort: 0, hostPort: 0 },
              rules: [{ required: true, message: '请输入IP设置' }],
            })(<ContainerNetworkInput style={{ width: '80%' }} />)}
          </Form.Item>
          <Form.Item label="端口填写说明">
            <Alert message="[ docker容器端口 ]: [ 对外暴露端口 ]" type="info" />
          </Form.Item>
        </div>
      </Form>
    );
  }
}

const WrappedContainerForm = Form.create({ name: 'containerForm' })(ContainerFormPanel);

class ContainerCreateModel extends React.Component {
  state = {
    form: null,
  };

  submit = e => {
    console.log(e);
    const { form } = this.state.form.props;
    form.validateFields((err, values) => {
      if (err) {
        return;
      }
      this.props.dispatch({
        type: 'dockerBasic/containerCreate',
        payload: {
          assetId: this.props.assetId,
          containerName: values.containerName,
          imageName: values.imageName,
        },
        callback: response => {
          if (response.Res) {
            message.success(`容器[${response.Obj.Id}]成功`);
            this.props.close();
          } else {
            message.error(`新增容器失败，原因:${response.Info}`);
          }
        },
      });
    });
  };

  render() {
    return (
      <Modal
        title="新增容器"
        visible={this.props.visible}
        onOk={this.submit}
        onCancel={this.props.close}
        confirmLoading={this.props.loading}
        destroyOnClose={true}
        afterClose={this.props.refreshInfo}
      >
        <WrappedContainerForm wrappedComponentRef={form => (this.state.form = form)} />
      </Modal>
    );
  }
}

function mapStateToProps(state) {
  return {
    loading: state.loading.models.dockerBasic,
  };
}

export default connect(mapStateToProps)(ContainerCreateModel);
