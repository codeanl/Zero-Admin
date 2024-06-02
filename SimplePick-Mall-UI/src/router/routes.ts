//对外暴露配置路由
export const constantRoute = [
    {
        path: '/login',
        component: () => import('@/view/login/index.vue'),
        name: 'login',
        meta: {
            title: 'login',
            hidden: true,
        },
    },
    {
        path: '/',
        component: () => import('@/view/layout/index.vue'),
        name: 'layout',
        meta: {
            title: '',
            hidden: false,
            icon: '',
        },
        redirect: '/home',
        children: [
            {
                path: '/home',
                component: () => import('@/view/home/index.vue'),
                meta: {
                    title: '首页',
                    hidden: false,
                    icon: 'HomeFilled',
                },
            },

        ],
    },
    {
        path: '/404',
        component: () => import('@/view/404/index.vue'),
        name: '404',
        meta: {
            title: '404',
            hidden: true,
        },
    },
    {
        path: '/setting',
        component: () => import('@/view/layout/index.vue'),
        name: 'Setting',
        meta: {
            title: '系统设置',
            hidden: false,
            icon: 'Setting',
        },
        redirect: '/setting/profile',
        children: [
            {
                path: '/setting/profile',
                component: () => import('@/view/setting/profile/index.vue'),
                name: 'Profile',
                meta: {
                    title: '个人设置',
                    icon: 'User',
                    hidden: false,
                },
            }, 
            {
                path: '/setting/setting',
                component: () => import('@/view/setting/setting/index.vue'),
                name: 'SystemSetting',
                meta: {
                    title: '系统设置',
                    icon: 'Setting',
                    hidden: false,
                },
            },
        ],
    },
    {
        path: '/merchantApply',
        component: () => import('@/view/pms/merchantApply/index.vue'),
        name: '商家入驻',
        meta: {
            title: 'merchantApply',
            hidden: true,
        },
    },
]

export const asyncRoute = [
    {
        path: '/sys',
        component: () => import('@/view/layout/index.vue'),
        name: 'Sys',
        meta: {
            title: '权限管理',
            hidden: false,
            icon: 'Lock',
        },
        redirect: '/sys/user',
        children: [
            {
                path: '/sys/user',
                component: () => import('@/view/sys/member/index.vue'),
                name: 'User',
                meta: {
                    title: '用户管理',
                    hidden: false,
                    icon: 'User',
                },
            },
            {
                path: '/sys/role',
                component: () => import('@/view/sys/role/index.vue'),
                name: 'Role',
                meta: {
                    title: '角色管理',
                    hidden: false,
                    icon: 'Avatar',
                },
            },
            {
                path: '/sys/menu',
                component: () => import('@/view/sys/menu/index.vue'),
                name: 'Menu',
                meta: {
                    title: '菜单管理',
                    hidden: false,
                    icon: 'List',
                },
            },
        ],
    },
    {
        path: '/pms',
        component: () => import('@/view/layout/index.vue'),
        name: 'Pms',
        meta: {
            title: '商品管理',
            hidden: false,
            icon: 'Goods',
        },
        redirect: '/pms/category',
        children: [
            {
                path: '/pms/category',
                component: () => import('@/view/pms/category/index.vue'),
                name: 'Category',
                meta: {
                    title: '分类管理',
                    icon: 'Management',
                    hidden: false,
                },
            },
            {
                path: '/pms/merchant',
                component: () => import('@/view/pms/merchant/index.vue'),
                name: 'Merchant',
                meta: {
                    title: '商家管理',
                    icon: 'Management',
                    hidden: false,
                },
            },
            {
                path: '/pms/attributeCategory',
                component: () => import('@/view/pms/attributeCategory/index.vue'),
                name: 'AttributeCategory',
                meta: {
                    title: '属性分类管理',
                    icon: 'CreditCard',
                    hidden: false,
                },
            },
            {
                path: '/pms/attribute',
                component: () => import('@/view/pms/attribute/index.vue'),
                name: 'Attribute',
                meta: {
                    title: '属性管理',
                    icon: 'Box',
                    hidden: false,
                },
            },
            {
                path: '/pms/product',
                component: () => import('@/view/pms/product/index.vue'),
                name: 'Product',
                meta: {
                    title: '商品管理',
                    icon: 'ShoppingCart',
                    hidden: false,
                },
            },
            {
                path: '/pms/place',
                component: () => import('@/view/pms/place/index.vue'),
                name: 'Place',
                meta: {
                    title: '自提点管理',
                    hidden: false,
                    icon: 'Position',
                },
            },
        ],
    },
    {
        path: '/log',
        component: () => import('@/view/layout/index.vue'),
        name: 'Log',
        meta: {
            title: '日志管理',
            hidden: false,
            icon: 'ChatSquare',
        },
        redirect: '/log/loginLog',
        children: [
            {
                path: '/log/loginLog',
                component: () => import('@/view/log/loginLog/index.vue'),
                name: 'LoginLog',
                meta: {
                    title: '登录日志',
                    icon: 'ChatLineSquare',
                    hidden: false,
                },
            },
            {
                path: '/log/sysLog',
                component: () => import('@/view/log/sysLog/index.vue'),
                name: 'SysLog',
                meta: {
                    title: '系统日志',
                    icon: 'ChatDotSquare',
                    hidden: false,
                },
            },
        ],
    },
    {
        path: '/ums',
        component: () => import('@/view/layout/index.vue'),
        name: 'Ums',
        meta: {
            title: '采购员管理',
            hidden: false,
            icon: 'User',
        },
        redirect: '/ums/member',
        children: [
            {
                path: '/ums/member',
                component: () => import('@/view/ums/member/index.vue'),
                name: 'Member',
                meta: {
                    title: '采购员管理',
                    icon: 'User',
                    hidden: false,
                },
            },

        ],
    },
    {
        path: '/sms',
        component: () => import('@/view/layout/index.vue'),
        name: 'Sms',
        meta: {
            title: '营销管理',
            hidden: false,
            icon: 'MessageBox',
        },
        redirect: '/sms/coupon',
        children: [
            {
                path: '/sms/homeAdvertise',
                component: () => import('@/view/sms/homeAdvertise/index.vue'),
                name: 'HomeAdvertise',
                meta: {
                    title: '广告管理',
                    icon: 'Discount',
                    hidden: false,
                },
            },
            {
                path: '/sms/coupon',
                component: () => import('@/view/sms/coupon/index.vue'),
                name: 'Coupon',
                meta: {
                    title: '优惠券管理',
                    icon: 'Ticket',
                    hidden: false,
                },
            },
            {
                path: '/sms/subject',
                component: () => import('@/view/sms/subject/index.vue'),
                name: 'Subject',
                meta: {
                    title: '专题管理',
                    icon: 'ShoppingCart',
                    hidden: false,
                },
            },
            {
                path: '/sms/subjectProduct',
                component: () => import('@/view/sms/subjectProduct/index.vue'),
                name: 'SubjectProduct',
                meta: {
                    title: '专题商品管理',
                    icon: 'Grid',
                    hidden: false,
                },
            },
        ],
    },
    {
        path: '/oms',
        component: () => import('@/view/layout/index.vue'),
        name: 'Oms',
        meta: {
            title: '订单管理',
            hidden: false,
            icon: 'ShoppingCart',
        },
        redirect: '/oms/order',
        children: [
            {
                path: '/oms/order',
                component: () => import('@/view/oms/order/index.vue'),
                name: 'Order',
                meta: {
                    title: '订单管理',
                    icon: 'ShoppingCart',
                    hidden: false,
                },
            },
            {
                path: '/oms/returnApply',
                component: () => import('@/view/oms/returnApply/index.vue'),
                name: 'ReturnApply',
                meta: {
                    title: '退货管理',
                    icon: 'Sell',
                    hidden: false,
                },
            },
            {
                path: '/oms/returnReason',
                component: () => import('@/view/oms/returnReason/index.vue'),
                name: 'ReturnReason',
                meta: {
                    title: '退货原因管理',
                    icon: 'SoldOut',
                    hidden: false,
                },
            },
        ],
    },

]

export const anyRoute = {
    path: '/:pathMatch(.*)*',
    redirect: '/404',
    name: 'Any',
    meta: {
        title: '任意路由',
        hidden: true,
    },
}
