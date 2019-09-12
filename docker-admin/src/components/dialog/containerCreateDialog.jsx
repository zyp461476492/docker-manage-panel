import React from 'react';
import { connect } from 'dva';
import { Modal, Form, Input, message } from 'antd';
import styles from './dialog.css';

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
          console.log(response);
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
