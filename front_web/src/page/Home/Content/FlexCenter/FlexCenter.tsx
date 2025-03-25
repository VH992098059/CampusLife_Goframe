import {Card, Empty, Flex, Skeleton, Typography} from 'antd';
import {FireOutlined} from "@ant-design/icons";
import { ActivityList } from "../../../../api/activity/activity";

import {useEffect, useState} from "react";
import {getUuid} from "../../../../router/token/token.tsx";
import {
    CardPhoneStyle,
    CardStyle, ImgPhoneStyle,
    ImgStyle
} from "../../../../utils/global/CardCss.ts";

/* 主页内容 */
const Remen = () => {
    const [items, setItems] = useState([]);
    const [page, setPage] = useState(1);
    const [size] = useState(8);
    const [isLoading, setIsLoading] = useState(true); // 添加加载状态
    const [hasMore, setHasMore] = useState(true); // 添加是否有更多数据的标志
    const [widthCard,setWidthCard] = useState(window.innerWidth);

    const GetCardId = (key: string) => {
        const url = getUuid() === null
          ? `/home/activity?key=${key}`
          : `/home/activity?key=${key}&uuid=${getUuid()}`;
        window.open(url, '_blank');
      };


    const fetchData = async (pageNum:number) => {
        try {
            setIsLoading(true);
            const res = await ActivityList(pageNum, size);
            if (res.list.length < size) {
                setHasMore(false);
            }
            if (pageNum === 1) {
                setItems(res.list);
            }
        } catch (err) {
            console.log(err);
        } finally {
            setIsLoading(false);
        }
    };
    useEffect(() => {
        const handleResize = () => {
            setWidthCard(window.innerWidth);
        };
        window.addEventListener('resize', handleResize);
        return () => {
            window.removeEventListener('resize', handleResize);
        };
    })
    useEffect(() => {
        fetchData(page);

    }, [page]);
    if(isLoading){
        return(
            <Skeleton active style={{padding:"0 30px"}}/>
        )
    }else if(items.length==0){
        return (
            <div style={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
            }}>
                <Empty
                    description="搜索未找到"

                />
            </div>
        )
    }
    else{
        return (
            <>
                {items.map((item) => (
                    /* eslint-disable-next-line @typescript-eslint/ban-ts-comment */
                    // @ts-expect-error
                    <Card hoverable style={widthCard>768?CardStyle:CardPhoneStyle} styles={{body: {padding: 0, overflow: 'hidden'}}} key={item.uuid} onClick={()=>{GetCardId(item.uuid)}}  >
                        <Flex justify="space-between">
                            <img
                                alt="avatar"
                                src="https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png"
                                style={widthCard>768?ImgStyle:ImgPhoneStyle}
                            />
                            <Flex vertical align="flex-end" justify="space-between" style={{padding: 20}}>
                                <Typography.Title level={5}>
                                    {/* eslint-disable-next-line @typescript-eslint/ban-ts-comment */}
                                    { // @ts-expect-error
                                        item.activity_title}
                                </Typography.Title>

                                <span
                                    id={"hot"}><FireOutlined/>&nbsp;{/* eslint-disable-next-line @typescript-eslint/ban-ts-comment */}
                                    { // @ts-expect-error
                                        item.popular}</span>
                            </Flex>
                        </Flex>
                    </Card>
                ))}
            </>
        );
    }
};
const FlexCenter = () => (
   <div>
       <Flex gap={"middle"} vertical={true}>
           <span style={{fontSize: "x-large", fontWeight: "bolder"}}>热门活动</span>
           <Flex gap="middle"  style={{flexWrap: 'wrap'}} className={"remen"} >
               {Remen()}
           </Flex>

       </Flex>

   </div>
);

export default FlexCenter;

