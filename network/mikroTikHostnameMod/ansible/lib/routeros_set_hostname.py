#!/usr/bin/env python3

from ansible.module_utils.basic import AnsibleModule
import librouteros
import librouteros.login

def run_module():
    module_args = dict(
        hostname=dict(type='str', required=True),
        host=dict(type='str', required=True),
        username=dict(type='str', required=True),
        password=dict(type='str', required=True, no_log=True),
    )

    result = dict(
        changed=False,
        original_message='',
        message='',
    )

    module = AnsibleModule(
        argument_spec=module_args,
        supports_check_mode=True
    )

    params = module.params
    hostname = params['hostname']
    host = params['host']
    username = params['username']
    password = params['password']

    try:
        api = librouteros.connect(username=username, password=password, host=host)
        identity = list(api(cmd='/system/identity/print'))[0]
        current_hostname = identity['name']
        api(cmd='/system/identity/set', **{'name': hostname})

        if current_hostname != hostname:
            if not module.check_mode:
                set_result = list(api(cmd='/system/identity/set', **{'name': hostname}))
            result['changed'] = True
            result['message'] = 'Hostname changed from {} to {} and set_result was {}'.format(current_hostname, hostname, set_result)

        else:
            result['message'] = 'Hostname is already set to {}'.format(hostname)

    except Exception as e:
        module.fail_json(msg=str(e))

    module.exit_json(**result)

def main():
    run_module()

if __name__ == '__main__':
    main()
